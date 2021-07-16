package controllers

import (
	"strconv"

	"github.com/RigottiG/go-books-api/database"
	"github.com/RigottiG/go-books-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't find book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "can't bind JSON " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't create book " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't find books" + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}
