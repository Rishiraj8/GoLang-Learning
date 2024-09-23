package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Gonna print the current time")
	presentTime := time.Now()

	fmt.Println(presentTime.Format("Monday Jan"))
	// format for printing what you need Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	//example we have to defaulty type monday to have the day printing
	fmt.Println("The current time only is : ", presentTime.Format("15:04:05"))
	createdTime := time.Date(2004, time.December, 8, 0, 0, 0, 0, time.Now().Location())
	fmt.Println(createdTime)
	//also tried to build a exe (a go application using these commands)
	//go mod init 05Time
	//GOOS=windows go build
	//for macs its GOOS="darwin"

}
