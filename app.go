package main

import (
	"fmt"
	"portfolio-api/delivery"
)

func main() {
	fmt.Println("Application Start")

	// Initialize server
	server := delivery.NewServer()
	delivery.Start(server)
}