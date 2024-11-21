package main

import (
	"fmt"
	"log"

	"github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Config"
	"github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Result"
	ResultReceiver "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Result"
)

func main() {
	// Load the configuration
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	// Start the result listener with a callback function
	ResultReceiver.StartResultListener(config.ServiceBus.Results, func(resultObj Result.OrderResult) {
		// Callback function that processes the received OrderResult
		fmt.Println("Received new result from service bus...")
		fmt.Println("Transaction ID:", resultObj.TransactionId)
	})
}
