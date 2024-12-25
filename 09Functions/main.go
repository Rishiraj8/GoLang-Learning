package main
import "fmt"

func greet( togreet string ) string{
	return "Hello " + togreet 
}  


func main()  {
	fmt.Println(greet("Rishi"))
}