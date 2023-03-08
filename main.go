package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//create gin api
	r := gin.Default()

	//create a new router group
	root := r.Group("/")

	//respond to GET requests on / with a HTML page
	root.GET("/", func(c *gin.Context) {
		const site string = index
		c.Data(200, "text/html; charset=utf-8", []byte(site))
	})

	root.GET("/timer", func(c *gin.Context) {
		const site string = timer
		c.Data(200, "text/html; charset=utf-8", []byte(site))
	})

	//listen and serve on port 4000
	r.Run(":4000")
}
