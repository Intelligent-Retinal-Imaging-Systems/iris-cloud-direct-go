package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Details for loaction of file in Azure Blob Storage
type Configuration struct {

	// Blob Storage Container
	ServiceBus ServiceBus

	// Azure File Name
	BlobStorage BlobStorage

	ClientGuid  string
	SiteLocalId string
}

type ServiceBus struct {
	Orders  string
	Results string
	Events  string
}

type BlobStorage struct {
	ImageSubmission string
	ContainerName   string
}

func LoadConfiguration(filePath string) (*Configuration, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open configuration file: %w", err)
	}
	defer file.Close()

	// Decode the JSON file
	var config Configuration
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode configuration JSON: %w", err)
	}

	return &config, nil
}
