package handler

import (
	"fmt"
	"magdalena-API/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Magdalena Tambunan",
		"bio":  "mahasiswi del",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Magdalena Tambunan",
		"bio":   "mahasiswi del",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")                                    // Mengambil parameter URL "id"
	title := c.Param("title")                              // Mengambil parameter URL "title"
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title}) // Mengembalikan JSON dengan id dan title
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")                                    // Mengambil parameter query "title"
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price}) // Mengembalikan JSON dengan title
}

func PostBooksHandler(c *gin.Context) {
	// data yang dimasukkan oleh user
	// menerima data title dan price
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput) // Menggunakan pointer
	if err != nil {

		errorMessages := []string{} // Perbaiki nama variabel dari errorMessage menjadi errorMessages
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) // Gunakan errorMessages yang sudah dideklarasikan
		}
		c.JSON(http.StatusBadRequest, gin.H{ // Perbaiki format JSON response
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{ // Menggunakan http.StatusOK
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
