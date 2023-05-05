package main

import "github.com/gin-gonic/gin"

func main() {
	greetings := "Hello"
	router := gin.Default()
	
	router.GET("/", func(context *gin.Context) {
//		context.HTML(200, "index.html", "")
//		context.File("index.html")
		context.Writer.WriteHeader(200)
		context.Writer.Write([]byte(greetings))
	})

	router.Run(":8181")
}
