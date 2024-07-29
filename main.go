package main

import(
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Working!!"})
	})

	r.Run(":8080")
}