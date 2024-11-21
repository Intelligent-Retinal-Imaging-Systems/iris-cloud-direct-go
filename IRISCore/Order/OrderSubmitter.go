package OrderSubmitter

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Request"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// Submit a new order to IRIS via service bus interface
func SubmitOrder(ServiceBusConnectionString string, OrderRequest Request.OrderRequest) error {
	json, err := SerializeOrderToJson(OrderRequest)

	if err != nil {
		return err
	}

	client := getServiceClient(ServiceBusConnectionString)

	sender, err := client.NewSender("orders", nil)
	if err != nil {
		panic(err)
	}
	defer sender.Close(context.TODO())
	sbMessage := &azservicebus.Message{
		Body: []byte(json),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)

	return err

}

// Convert OrderRequest object to JSON for submission to the service bus
func SerializeOrderToJson(OrderRequest Request.OrderRequest) (string, error) {

	jsonData, err := json.Marshal(OrderRequest)

	return string(jsonData), err
}

func getServiceClient(connectionString string) *azservicebus.Client {
	// Create a new service client with connection string
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	handleError(err)

	return client
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
