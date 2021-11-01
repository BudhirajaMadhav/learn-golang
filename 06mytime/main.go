package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()

	fmt.Println(presentTime)
	// Output: 2021-11-01 22:12:57.980470596 +0530 IST m=+0.000092515

	fmt.Println(presentTime.Format("2/1/2006, Mon 15:04.0000 -0700 MST"))
	// Output: 1/11/2021, Mon 22:12.1008 +0530 IST

	/*
			Format returns a textual representation of the time value formatted
		according to layout, which defines the format by showing how the reference
		time, defined to be
			Mon Jan 2 15:04:05 -0700 MST 2006
		would be displayed if it were the value; it serves as an example of the
		desired output. The same display rules will then be applied to the time
		value.

		A fractional second is represented by adding a period and zeros
		to the end of the seconds section of layout string, as in "15:04:05.000"
		to format a time stamp with millisecond precision.
	*/

	createdTime := time.Date(2022, time.February, 19, 2, 30, 00, 00, time.Local)
	fmt.Println(createdTime.Format("Monday 2-1-2006 15:04 MST"))

}
