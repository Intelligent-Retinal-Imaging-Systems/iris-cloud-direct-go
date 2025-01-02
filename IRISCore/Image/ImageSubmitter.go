package ImageSubmitter

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func getServiceClientFromConnectionString(connectionString string) *azblob.Client {
	if connectionString == "" {
		log.Fatal("Connection string is empty")
	}

	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("Failed to create azblob client from connection string: %v", err)
	}

	return client
}

// Submit a new image to IRIS using a file on local machine
func UploadBlobFileFromLocalPath(connectionString string, containerName string, blobName string, path string) error {

	var client = getServiceClientFromConnectionString(connectionString)
	// Open the file for reading
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	handleError(err)

	defer file.Close()

	// Upload the file to the specified container with the specified blob name
	_, err = client.UploadFile(context.TODO(), containerName, blobName, file, nil)
	handleError(err)

	return err
}

// Submit a new image to IRIS using a stream
func UploadBlobAsStream(connectionString string, containerName string, blobName string, data string) error {

	var client = getServiceClientFromConnectionString(connectionString)
	blobContentReader := strings.NewReader(data)

	// Upload the file to the specified container with the specified blob name
	_, err := client.UploadStream(context.TODO(), containerName, blobName, blobContentReader, nil)
	handleError(err)

	return err
}

// Submit a new image to IRIS using binary data
func UploadBlobFromBinary(connectionString string, containerName string, blobName string, data []byte) error {

	var client = getServiceClientFromConnectionString(connectionString)
	// Upload the data to a block blob
	_, err := client.UploadBuffer(context.TODO(), containerName, blobName, data, nil)
	handleError(err)

	return err
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
