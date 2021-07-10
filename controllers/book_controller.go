package controllers

import (
	"strconv"

	"github.com/PauloDavi/webapi-with-go/database"
	"github.com/PauloDavi/webapi-with-go/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Id has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can not find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can not create book: " + err.Error(),
		})

		return
	}

	c.JSON(201, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can not list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can not create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeteleBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Id has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error": "Can not delete book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}