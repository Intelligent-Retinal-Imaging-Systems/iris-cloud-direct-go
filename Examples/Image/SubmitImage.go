package main

import (
	"log"
	"os"

	ImageSubmitter "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Image"
	"github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Config"
)

func main() {
	//SubmitFromLocal()
	//SubmitFromStream()
	SubmitFromBinary()
}

// Submit image from local file
func SubmitFromLocal() {
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	var filePath = "C:\\Users\\jholcomb\\Downloads\\2024-12-03_143634438.dcm"

	err = ImageSubmitter.UploadBlobFileFromLocalPath(config.BlobStorage.ImageSubmission, config.BlobStorage.ContainerName, "sample.dcm", filePath)
	if err != nil {
		log.Fatalf("Failed to submit image: %v", err)
	} else {
		log.Println("Image successfully submitted")
	}
}

func SubmitFromStream() {
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	var filePath = "C:\\Users\\jholcomb\\Downloads\\2024-12-03_143634438.dcm"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to get path: %v", err)
	}

	data := string(fileContent)

	err = ImageSubmitter.UploadBlobAsStream(config.BlobStorage.ImageSubmission, config.BlobStorage.ContainerName, "sampleFromStream.dcm", data)
	if err != nil {
		log.Fatalf("Failed to submit image: %v", err)
	} else {
		log.Println("Image successfully submitted")
	}
}

func SubmitFromBinary() {

	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	var filePath = "C:\\Users\\jholcomb\\Downloads\\2024-12-03_143634438.dcm"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to get path: %v", err)
	}

	data := []byte(fileContent)

	err = ImageSubmitter.UploadBlobFromBinary(config.BlobStorage.ImageSubmission, config.BlobStorage.ContainerName, "sampleFromBinary.dcm", data)
	if err != nil {
		log.Fatalf("Failed to submit image: %v", err)
	} else {
		log.Println("Image successfully submitted")
	}
}
