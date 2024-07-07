package main

import (
	"strconv"
	"time"
	"errors"
	"fmt"
)

type Cuisine int

const(
	CUISINE_UNSPECIFIED Cuisine = 1
	CUISINE_NORTH_INDIAN Cuisine = 1
	CUISINE_SOUTH_INDIAN Cuisine = 2
)

type Restraunt struct {
	id string
	name       string
	address    *Address
	costForTwo int
	capacity   int
	openTiming int
	closeTime int
	cuisine Cuisine
	bookingBuffer int
	dateToBookingMap map[string]DayBookingInterface
}

func getDateString(tm time.Time) string {
	year, month, day := tm.Date()
	 return strconv.Itoa(year) + month.String() + strconv.Itoa(day) 
}

func NewRestraunt(id string, name string, address *Address, costForTwo, capacity, openTime, 
	closeTime, bookingBuffer int, cuisine Cuisine) *Restraunt{

		dateToBookingMap := make(map[string]DayBookingInterface)
		
		for i:=0; i<bookingBuffer; i++{
			tm := time.Now().AddDate(0, 0, i)
			dateString := getDateString(tm)

			dayBooking := NewDayBooking(capacity, openTime, closeTime)
			dateToBookingMap[dateString] = dayBooking
		}

		return &Restraunt{
			id: id,
			name: name,
			address: address,
			costForTwo: costForTwo,
			openTiming: openTime,
			closeTime: closeTime,
			cuisine: cuisine,
			bookingBuffer: bookingBuffer,
			dateToBookingMap: dateToBookingMap,
		}	
}
 

func (r *Restraunt) GetName() string {
	if r == nil {
		return ""
	}

	return r.name
}

func (r *Restraunt) GetId() string {
	if r == nil {
		return ""
	}

	return r.id
}

func (r *Restraunt) GetAddress() *Address {
	if r == nil {
		return nil
	}

	return r.address
}

func (r *Restraunt) GetCostForTwo() int {
	if r == nil {
		return 0
	}

	return r.costForTwo
}

func (r *Restraunt) GetCuisine() Cuisine {
	if r == nil {
		return CUISINE_UNSPECIFIED
	}

	return r.cuisine
}

func (r *Restraunt) IsBookingAvailable(date string, time, numPeople int) (bool, error) {
	dayBooking, ok := r.dateToBookingMap[date]
	if !ok {
		return false, errors.New("cant book for this date")
	}

	ok, err := dayBooking.IsBookingAvailable(time, numPeople)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (r *Restraunt) BookSlot(date string, time, numPeople int) error {
	dayBooking, ok := r.dateToBookingMap[date]
	if !ok {
		return errors.New("cant book for this date")
	}

	bookingErr := dayBooking.BookSlot(time, numPeople)
	if bookingErr != nil {
		return bookingErr
	}

	fmt.Printf("slot booked for %v people at time %v at restraunt: %v", numPeople, time, r.GetName())
	fmt.Println()

	return nil
}


