// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

func main() {
	// ------------------------------------------------------------
	// Create nil slices
	// ------------------------------------------------------------

	// Pizza toppings
	var pizza []string

	// Departure times
	var departures []time.Time

	// Student graduation years
	var grads []int

	// On/off states of lights in a room
	var lights []bool

	// ------------------------------------------------------------
	// Append them some elements (use your creativity!)
	// ------------------------------------------------------------
	pizza = append(pizza, "pepperoni", "onions", "extra cheese")

	now := time.Now()
	departures = append(departures,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	grads = append(grads, 1998, 2005, 2018)

	lights = append(lights, true, false, true)

	// ------------------------------------------------------------
	// Print the slices
	// ------------------------------------------------------------

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ngraduations : %d\n", grads)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\nlights      : %t\n", lights)
}