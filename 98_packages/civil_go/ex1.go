package main

import (
	"fmt"
	"time"

	"cloud.google.com/go/civil"
)


func main() {



	civilDate := civil.DateOf(time.Now())
	fmt.Printf("civilDate: %v %T\n", civilDate, civilDate)
	
	myDate := civilDate.String()
	fmt.Printf("myDate: %v %T\n", myDate, myDate)



	// Add days
	birthday,_ := civil.ParseDate("1996-08-26")
	fmt.Printf("dt: %v\n", birthday)
	// lets add
	// birthday = birthday.AddDays(10000)
	fmt.Printf("birthday: %v\n", birthday)


	// Days since my birthday
	daysSince := civilDate.DaysSince(birthday) // DaysSince returns the signed number of days between the date and s, not including the end day. This is the inverse operation to AddDays.

	fmt.Printf("Days in Earth completed : %v\n", daysSince)
	

}