package main

import (
	"app/database"

	"github.com/gin-gonic/gin"
	"log"
)

func main(){

	database.SetupDB()

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Working!!"})
	})

	err := r.Run(":8080")
	if err!=nil{
		log.Fatal(err)
	}
}