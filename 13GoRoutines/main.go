package main

import (
	"fmt"
	"time"
)

// Go Routines
// Go routines are lightweight threads of execution.
// They are used to execute functions concurrently.
// Go routines are created using the `go` keyword followed by a function invocation.
// The function runs concurrently in a new go routine.
// The main function is also a go routine.
// The main function is the first go routine that runs in a Go program.
func printNumber(number int) {
	fmt.Println(number);
}
func main()  {
	go printNumber(1);
	go printNumber(2);
	go printNumber(3);
	go printNumber(4);
	go printNumber(5);
	
	go printNumber(6);
	go printNumber(7);
	go printNumber(8);
	go printNumber(9);
	go printNumber(10);
	fmt.Println("gets printed first");
	time.Sleep(time.Second*2);
	fmt.Println("gets printed last");
	//numbers will be printed in random order
	//because the go routines are running concurrently
	fmt.Print()
}
