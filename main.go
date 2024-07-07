package main

import (
	"fmt"
	"time"
)

// RegisterRestaurant : Any restaurant owner should be able to come on the platform and register a restaurant.
// Restaurants should be able to also provide the slots when it is available for bookings,

// SearchRestaurant : Any logged in user, should be able to search a restaurant based on city, area, cuisine, restaurantName,
// cost for two, veg/non-veg etc

// BookTable : Once the user is done searching the restaurant,
// he/she should be able to book a table for n people in the
// restaurant if the table and slot is available for the booking date.
// For simplicity, you can assume all tables in the restaurant are eligible for n people and only 1 hour time slots are allowed for booking.
// Bookings only allowed for upto m days in future

// Register Restraunt - (name, address, cost for two, capacity, timings, cuisine, veg-non/veg) (error)
// SearchRestraunt - (req{}) - list of restraunts
// GetSlots - (restrauntId, date, numPeople) list of slots
// BookTable(restrauntId, startTime, numPeople) (error)

// Restraunt
// Slot
// Booking
// BookingForTheDay

func main (){
	rest1 := NewRestraunt("rest1", "pizza hut",&Address{
		city:"Bangalore",
	}, 1000, 10, 10, 22, 5, CUISINE_NORTH_INDIAN)
	rest2 := NewRestraunt("rest2", "kfv", &Address{
		city:"Bangalore",
	}, 1000, 10, 10, 22, 5, CUISINE_NORTH_INDIAN)
	rest3 := NewRestraunt("rest3", "mcd", &Address{
		city:"Delhi",
	}, 2000, 10, 10, 22, 5, CUISINE_SOUTH_INDIAN)

	restList := []*Restraunt{}

	sliceStorage := &RestrauntStorageSlice{
		restList,
	}

	restManager := RestrauntManger{
		storage: sliceStorage ,
	}

	// register calls
	restManager.RegisterRestraunt(rest1)
	restManager.RegisterRestraunt(rest2)
	restManager.RegisterRestraunt(rest3)

	// search for restraunts
	specs := []RestrauntSpecification{}
	specs = append(specs, &RestrauntSpecificationCity{city: "Bangalore"})
	specs = append(specs, &RestrauntSpecificationCuisine{cuisuin: CUISINE_NORTH_INDIAN})

	res, err := restManager.SearchRestraunt(specs)
	if err != nil{
		fmt.Println(fmt.Sprintf("error in search: %v", err))
	}
	fmt.Println(fmt.Sprintf("search result: %v", res))


	dateStr := getDateString(time.Now())
	// book restraunts
	bookErr := restManager.BookTable("rest1", dateStr, 15, 9)
	if bookErr != nil{
		fmt.Println(fmt.Sprintf("error in BookTable: %v", bookErr))
	}

	// try to overbook
	bookErr = restManager.BookTable("rest1", dateStr, 15, 9)
	if bookErr != nil{
		fmt.Println(fmt.Sprintf("error in BookTable: %v", bookErr))
	}
}