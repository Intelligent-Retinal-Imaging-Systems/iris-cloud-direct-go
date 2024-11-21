package ResultReceiver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Result"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// Start listener with connection string and callback function for new results
func StartResultListener(ServiceBusConnectionString string, cb func(Result.OrderResult)) {
	client := getServiceClient(ServiceBusConnectionString)

	receiver, err := client.NewReceiverForQueue("results", nil)
	if err != nil {
		log.Fatalf("Failed to create receiver: %v", err)
	}
	defer receiver.Close(context.TODO())

	for {
		// Receive a batch of messages (increase the max number of messages to receive more at once)
		messages, err := receiver.ReceiveMessages(context.TODO(), 10, nil) // Adjust 10 based on your requirements
		if err != nil {
			log.Printf("Failed to receive messages: %v", err)
			continue // Skip and retry on failure
		}

		// Check if there are no messages to process
		if len(messages) == 0 {
			log.Println("No more messages in the queue.")
			break
		}

		for _, message := range messages {
			// Process each received message
			body := message.Body

			// Attempt to deserialize the message body to a result object
			resultObj, err := GetResultFromJsonSafe(string(body)) // Use a safe deserialization function
			if err != nil {
				// Log the error but continue processing other messages
				log.Printf("Failed to process message: %v. Skipping this message.", err)
				// Optionally, you can complete the message even if it's bad, or just leave it in the queue
				continue
			}

			// Call the provided callback with the result object
			cb(resultObj)

			// Complete the message to remove it from the queue
			err = receiver.CompleteMessage(context.TODO(), message, nil)
			if err != nil {
				log.Printf("Failed to complete message: %v", err)
			}
		}
	}
}

// A safer version of GetResultFromJson that returns an error if deserialization fails
func GetResultFromJsonSafe(jsonStr string) (Result.OrderResult, error) {
	var resultObj Result.OrderResult
	err := json.Unmarshal([]byte(jsonStr), &resultObj)
	if err != nil {
		return resultObj, fmt.Errorf("failed to deserialize message: %w", err)
	}
	return resultObj, nil
}

func getServiceClient(connectionString string) *azservicebus.Client {
	// Create a new service client with connection string
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create service client: %v", err)
	}
	return client
}
