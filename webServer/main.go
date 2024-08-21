package main

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func putAlbums(context *gin.Context) {
	var newAlbum album
	if err := context.BindJson(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/allAlbums", getAlbums)
	router.POST()
	router.Run("127.0.0.1:8080")
}
