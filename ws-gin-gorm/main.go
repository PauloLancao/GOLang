package main

import (
	"net/http"

	"github.com/PauloLancao/GOLang/ws-gin-gorm/controllers"
	"github.com/PauloLancao/GOLang/ws-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	// routes
	router.GET("/", hello)
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}
