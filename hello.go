package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	
	router.GET("/", func(context *gin.Context) {
//		context.HTML(200, "index.html", "")
		context.File("index.html")
	})

	router.Run(":8181")
}
