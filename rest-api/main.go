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
func getStudentByRegno(c *gin.Context) {
	regno := c.Param("regno") // Get URL parameter
	c.String(http.StatusOK, "Requested Reg No: %s", regno)
}
func getStudents(c *gin.Context){
	c.IndentedJSON(http.StatusOK,students);
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
func searchStudents(c *gin.Context) {
	dept := c.Query("dept") // Extract "dept" query param
	c.String(http.StatusOK, "Searching for students in %s department", dept)
}
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums",postAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.GET("/students",getStudents);
    router.GET("/students/:regno", getStudentByRegno)
    router.GET("/search", searchStudents)
    router.Run("localhost:8080")
}


