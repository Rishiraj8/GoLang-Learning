package main

import (
	"fmt"
	"net/http"
	"log"
)


func main(){
	http.Handle("/",http.FileServer(http.Dir("16go-server/static")))
	port := ":5080"
	http.HandleFunc("/about",
	func(w http.ResponseWriter,r *http.Request){})


	
	fmt.Println("server listening on port"+port)
	log.Fatal(http.ListenAndServe(port,nil))
}