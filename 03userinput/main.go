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
