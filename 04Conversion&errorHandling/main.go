package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Just Type your phone number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// convert, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	// // the above code will work since we have trimmed the unwanted \r and \n came from reader
	convert, err := strconv.ParseInt(input, 10, 64)
	// for string it needs to get three parameters
	// one is input another is (type:decimal,binary,hexa: in this case its decimal so its 10)
	// third one is the size of the int (uint64)
	if err != nil {
		fmt.Println("This is the error you are encountering", err)
	} else {
		fmt.Println("This is the your phone number in india is +91", convert)
	}
}
