package main

import (
	EventReceiver "iris-cloud-direct-go/IRISCore/Event"
	"iris-cloud-direct-go/IRISCore/Models/Config"

	"fmt"
	"log"
)

func main() {
	// Load the configuration
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	// Start the result listener with a callback function
	EventReceiver.StartEventListener(config.ServiceBus.Results, func(msg string) {
		fmt.Println("Received new result from service bus...")
		fmt.Println("Transaction ID:", msg)
	})
}
