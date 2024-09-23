package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)
	// //will give <NIl> since we haven't intialized any pointers
	myNumber := 8
	ptr := &myNumber
	//or var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr * 2
	fmt.Println(myNumber)
}
