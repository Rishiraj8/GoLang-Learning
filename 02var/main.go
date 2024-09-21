package main

import "fmt"

//jwttoken := 300000
//it gives error because walrus operator is only used
//inside any method only allowed to use
//but outside the method is not allowed
var jwttoken = 300000

//allowed
var jwttoken2 int = 30000

//allowed
const LogicToken string = "givasdel;j" //public

// if the first letter is "capital" then its global variable
const logictoken string = "asdf"

// this is not a public token

func main() {
	fmt.Print("main.go\n")
	var rishi string = "rishi"
	fmt.Printf("variable_type %T\n", rishi)
	var login bool = true
	fmt.Println(login)
	var small_num uint8 = 255
	//int is most commanly used
	fmt.Println(small_num)
	var float_num float64 = 1234.1234567890
	fmt.Println(float_num)
	var num int
	fmt.Println(num)
	var string1 string
	fmt.Println(string1)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	// not a good way of doing without describing the type
	// website = 3
	// fmt.Println(website)
	//02var\main.go:25:12: cannot use 3 (untyped int constant) as string value in assignment

	numberofUser := 100000.0123
	testtype := 10000.0
	//walrus operator
	fmt.Printf("testtype number is %f and %T\n", testtype, testtype)
	fmt.Println(testtype)
	//just tried variety of these
	fmt.Println(numberofUser)
	fmt.Println(LogicToken)
}
