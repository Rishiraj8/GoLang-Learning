package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	var names [4]string
	names[0] = "Rishi"
	names[1] = "Sarath"
	names[2] = "Rajvikash"
	fmt.Println(names)
	fmt.Println(len(names))
	var rollno = [10]int{137, 149, 133}
	fmt.Println(rollno)
	fmt.Println(len(rollno))
}
