package main

import "fmt"


func main() {
	students := map[string]int{
		"Rishi": 137,
		"Sarath": 149,
		"Rajvikash": 133,
	}
	fmt.Println(students)
	badstudent := "Rishi"
	dirty,exists := students[badstudent]
	if exists {
		fmt.Println(dirty)
	} else {
		fmt.Println("Student not found")
	}	

}