package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name         string
	Age          int
	IsInterested bool
}

func sayHello(p *Person) string {
	return "Hello" + p.Name
}

func main() {

	p := Person{Name: "john", Age: 20, IsInterested: true}

	s := sayHello(&p)

	fmt.Printf(s)
	fmt.Printf(p.Name)

	//fmt.Println("Hello World")
	r := gin.Default()
	// Load HTML files. Use LoadHTMLGlob if you have multiple HTML files.
	r.LoadHTMLFiles("./frontend/index.html")

	// Set up a route to serve the HTML file.
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	//post a person to the server
	//r.POST("/person", func())

	// Serve static files. The URL path is "/static" and files are served from the "./frontend" directory.
	// Adjust the paths according to your directory structure.
	r.Static("/static", "./frontend")

	// Start the server on port 8001
	r.Run(":8001")
	r.Static("/static", "./frontend")

	// Start the server on port 8001
	r.Run(":8001")

}
