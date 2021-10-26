package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()

	// Format MM-DD-YYYY
	fmt.Println(dt.Format("01-02-2006"))

	// Format MM-DD-YYYY hh:mm:ss
	fmt.Println(dt.Format("01-02-2006 15:04:05"))

	// With short weekday (Mon)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Mon"))

	// With weekday (Monday)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))

	// Include micro seconds
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000"))

	timeNow := time.Now()
	fmt.Println("The Present time is ", timeNow)
	timeVar := time.Date(2021, time.October, 11, 22, 1, 9, 9, time.UTC)
	fmt.Println("The Defined time is ", timeVar)
	fmt.Println("The Defined time is ", timeVar.Format(time.ANSIC))
	timeParsed, _ := time.Parse(time.ANSIC, "Mon Oct 11 22:01:09 2021")
	fmt.Printf("The type of timeParsed variable is %T\n", timeParsed)

	hours, minutes, _ := time.Now().Clock()
	currUTCTimeInString := fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	hours, minutes = 9, 0
	currUTCTimeInString = fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)
	hours, minutes = 10, 0
	currUTCTimeInString = fmt.Sprintf("%d%02d", hours, minutes)
	fmt.Println(currUTCTimeInString)

}
