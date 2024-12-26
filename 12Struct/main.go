package main

import "fmt"

//note that the struct doesn't any commons between the fields
type book struct {
	title string
	author string
	pages int
}
// Method associated with the `book` type
//see the difference between the function and method
//this is method
func (b book) details() {
    fmt.Printf("'%s' by %s, %d pages\n", b.title, b.author, b.pages)
}

//this is function
func printDetails(title string, author string, pages int) {
    fmt.Printf("'%s' by %s, %d pages\n", title, author, pages)
}
// Why Use Methods?
// Encapsulation: Methods provide a cleaner, more intuitive way to operate on a type and its fields.
// Reusability: You don't need to pass the fields of the type explicitly—methods can directly access them.
// Readability: Code like book1.details() is more natural and readable than calling a standalone function with arguments.

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
	book1.details()
	printDetails(book2.title, book2.author, book2.pages)
}