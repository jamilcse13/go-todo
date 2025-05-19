// main.go
package main

import (
	"go-todo/config"
	"go-todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()

	r := gin.Default()

	// Initialize router
	routes.SetupRouter(r)

	// run
	port := ":9090"
	r.Run(port)
}
