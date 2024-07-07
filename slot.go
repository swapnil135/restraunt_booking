package main

type Slot struct {
	startTime int
	endTime int
	isBooked bool
}

func NewSlot(startTime, endTime int) *Slot{
	return &Slot{
		startTime: startTime,
		endTime: endTime ,
		isBooked: false,
	}
}

func(s *Slot) GetStartTime() int {
	if(s == nil) {
		return -1
	}

	return s.startTime
}

func(s *Slot) GetEndTime() int {
	if(s == nil) {
		return -1
	}

	return s.endTime
}

func(s *Slot) IsBooked() bool {
	if(s == nil) {
		return false
	}

	return s.isBooked
}