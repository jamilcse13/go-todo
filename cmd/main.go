// main.go
package main

import (
	"fmt"

	"github.com/jamilcse13/go-todo/config" // Import config package
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()

	// Initialize router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go API with PostgreSQL"})
	})

	// Start server
	port := ":9090"
	fmt.Println("Server is running on port", port)
	r.Run(port)
}
