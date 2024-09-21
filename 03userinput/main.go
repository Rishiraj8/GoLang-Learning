package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Just try to give me your email.ID")
	reader := bufio.NewReader(os.Stdin)
	//bufio is a pacage where New Reader is a function;
	//os is package stdin is function
	//New Reader is used to make a reader function
	//stdin is used to make a input/output for activating the reading function
	//poi docs ahh pathalum onnum pureyathu
	//you can also use C syntax just tried it with these
	//https://youtu.be/zYIZtbyUIDY?si=JNFK2hZhkuR2ssDn
	//for tutorial
	email, _ := reader.ReadString('\n')
	//ReadLine is used to input as a string,
	// get the input until '\n'
	// '_' in email,_ is used if you are trying to do (try,catch)
	// (where try is email) and (catch is _) professionally done as _ you are not using as a try and catch
	fmt.Println("This is your email Id", email)
	fmt.Printf("This is the type of email %T", email)

}

// package main
// import ("fmt")

// func main() {
//   var i,j int

//   fmt.Print("Type two numbers: ")
//   fmt.Scanf("%v %v",&i, &j)
//   fmt.Println("Your numbers are:", i, "and", j)
// }
//Note: %v tells Scanf() to store the value of the inputs in the variables.
//In this example, the inputs are received in exactly the same way defined in Scanf() formatting.

// This means that the code expects inputs to be received with one space between them. Like: 1 2.

// package main
// import ("fmt")

// func main() {
//   var i,j int

//   fmt.Print("Type two numbers: ")
//   fmt.Scanln(&i, &j)
//   fmt.Println("Your numbers are:", i, "and", j)
// }
