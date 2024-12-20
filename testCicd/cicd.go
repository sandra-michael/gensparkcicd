package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //  Default returns an Engine instance with the Logger and Recovery middleware already attached.

	// Route Parameters
	// name is going to be a parameter
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, you pinged !!!!!!!!")
	})

	err := r.Run(":8083")
	if err != nil {
		panic(err)
	}

}
