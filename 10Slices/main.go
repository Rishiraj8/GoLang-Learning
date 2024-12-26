package main

import "fmt"

func main() {
	fmt.Println("Slices in Go")
	var names = []string{"Rishi", "Sarath", "Rajvikash"}
	fmt.Println(names)
	fmt.Println(len(names))
	rollno := []int{137, 149, 133}
	fmt.Println(rollno)
	fmt.Println(len(rollno))
	//see this importantly how to append
	names = append(names, "Yogeshwaran")
	fmt.Println(names)
	names = append(names, "Sarath", "Rajvikash")
	fmt.Println(names)
}