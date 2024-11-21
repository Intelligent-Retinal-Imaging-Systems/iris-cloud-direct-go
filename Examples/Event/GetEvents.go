package main

import (
	"fmt"
	EventReceiver "iris-cloud-direct-go/IRISCore/Event"
	"iris-cloud-direct-go/IRISCore/Models/Config"
	"log"
)

func main() {
	// Load the configuration
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	// Start the event listener with a callback function
	EventReceiver.StartEventListener(config.ServiceBus.Results, func(resultObj string) {
		// Callback function that processes the received OrderResult
		fmt.Println("Received new result from service bus...")
		fmt.Println("Transaction ID:", resultObj)
	})
}
