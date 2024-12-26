package main

import "fmt"

//note that the struct doesn't any commons between the fields
type book struct {
	title string
	author string
	pages int
}
func main()  {
	//but here we have to follow the order of the struct
	// and also put commos between the fields
	book1 := book{"Harry Potter", "JK Rowling", 300}
	fmt.Println(book1)
	book2 := book{
		pages:  200,
        title:  "The Alchemist",
        author: "Paulo Coelho",
    }
	fmt.Println(book2)
}