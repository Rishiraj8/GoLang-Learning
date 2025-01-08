package main

import (
	"fmt"
	"net/http"
	"log"

)		
func handleabout(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/about"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w,"Welcome to the about page")
}

func handlehello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w,"Hello, World!")
}

func main()  {
	http.Handle("/",http.FileServer((http.Dir("17proper-structured-http-server/static"))))


    http.HandleFunc("/about",handleabout)
	http.HandleFunc("/hello",handlehello)
	port := ":5080"
	log.Fatal(http.ListenAndServe(port,nil))
}