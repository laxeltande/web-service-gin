package main

import (
	"github.com/gin-gonic/gin"
	//go get . in command line to begin tracking the Gin module as a dependency.
	"net/http"
)

// album represents data about a record album. You’ll use this to store album data in memory.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// assign the handler function to an endpoint path
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums) //Why do we pass the name of the function??
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// albums slice to seed record album data. (containing data you’ll use to start)
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// creates JSON from the slice of album structs, writing the JSON into the response.
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
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
