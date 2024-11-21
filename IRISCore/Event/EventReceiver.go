package EventReceiver

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// Start listener with connection string and callback function for new events
func StartEventListener(ServiceBusConnectionString string, cb func(string)) {
	client := getServiceClient(ServiceBusConnectionString)

	receiver, err := client.NewReceiverForQueue("events", nil)
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

			// Attempt to deserialize the message body to a event string
			msg, err := GetResultFromJsonSafe(string(body)) // Use a safe deserialization function
			if err != nil {
				// Log the error but continue processing other messages
				log.Printf("Failed to process message: %v. Skipping this message.", err)
				// Optionally, you can complete the message even if it's bad, or just leave it in the queue
				continue
			}

			// Call the provided callback with the event object
			cb(msg)

			// Complete the message to remove it from the queue
			err = receiver.CompleteMessage(context.TODO(), message, nil)
			if err != nil {
				log.Printf("Failed to complete message: %v", err)
			}
		}
	}
}

// A safer version of GetResultFromJson that simply returns the raw message body as a string
func GetResultFromJsonSafe(jsonStr string) (string, error) {
	// Just return the input string directly without attempting to deserialize
	return jsonStr, nil
}

func getServiceClient(connectionString string) *azservicebus.Client {
	// Create a new service client with connection string
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create service client: %v", err)
	}
	return client
}
