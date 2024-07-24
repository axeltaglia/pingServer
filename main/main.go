package main

import (
	"basic_http_backend_service/server"
	"fmt"
	"log/slog"
)

func main() {
	// Initialize the server on port 8081
	api := server.NewServer(":8081")
	// Set up the endpoints for the server
	api.HandleEndpoints()

	fmt.Println("Server running at port 8081")
	// Run the server and log an error if it fails to start
	if err := api.Run(); err != nil {
		slog.Error("Server couldn't start: ", err)
	}
}
