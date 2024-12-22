package main

import (
	"github.com/gin-gonic/gin"
	// Berikan alias jika menggunakan goccy/json
	"magdalena-API/handler"
)

func main() {
	// Membuat router default dengan Gin
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.GET("/books", handler.PostBooksHandler)

	router.Run(":8000")
}
