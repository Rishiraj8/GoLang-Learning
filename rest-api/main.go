package main

import (
        "net/http"
        "github.com/gin-gonic/gin"
)

// album represents data about a record album.
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
type student struct{
	Regno int `json:"reg_no"`
	Name string `json:"name"`
	Dept string `json:"dept"`
	Year int `json:year`
}

var students = []student{
	{Regno: 722822104137, Name: "Rishi", Dept:"CSE",Year: 3},
	{Regno: 722822104134, Name: "RajVikash", Dept:"CSE",Year: 3},
	{Regno: 722822104149, Name: "Sarath", Dept:"CSE",Year: 3},
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
func getStudents(c *gin.Context){
	c.IndentedJSON(http.StatusOK,students);
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}


func main() {
    router := gin.Default()
    router.GET("/", getAlbums)
    router.GET("/students",getStudents);

    router.Run("localhost:8080")
}


