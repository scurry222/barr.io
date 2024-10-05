// main.go
package main

import (
	"log"

	"barr.io/api/routes"
	"barr.io/db"
)

func main() {
	// Initialize the database
	db.ConnectDatabase()

	// Set up the router
	router := routes.SetupRouter()

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
