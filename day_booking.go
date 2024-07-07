package main

import "errors"
import "fmt"

// takes care of all booking for a day for a resource
type DayBookingInterface interface {
	IsBookingAvailable(time, numPeople int) (bool, error)
	BookSlot(time, numPeople int) error
}
type DayBooking struct {
	capacity         int
	openTime         int
	closeTime        int
	bookingToHourMap map[int]int
}

var _ DayBookingInterface = &DayBooking{}

func NewDayBooking(capacity, openTime, closeTime int) *DayBooking {
	bookingMap := map[int]int{}
	for i := openTime; i < closeTime; i++ {
		bookingMap[i] = capacity
	}

	return &DayBooking{
		capacity:         capacity,
		openTime:         openTime,
		closeTime:        closeTime,
		bookingToHourMap: bookingMap,
	}
}

func (d *DayBooking) IsBookingAvailable(time, numPeople int) (bool, error) {
	avl, ok := d.bookingToHourMap[time]
	if !ok {
		return false, errors.New(fmt.Sprintf("booking outside working hours, for time: %v", time))
	}

	return avl>= numPeople, nil
}

func (d *DayBooking) BookSlot(time, numPeople int) error {
	ok, err := d.IsBookingAvailable(time, numPeople);
	if err != nil {
		return errors.Join(err, errors.New("error in is booking avl"))
	}

	if !ok {
		return errors.New("slot not available")
	}

	d.bookingToHourMap[time]-=numPeople

	

	return nil
}