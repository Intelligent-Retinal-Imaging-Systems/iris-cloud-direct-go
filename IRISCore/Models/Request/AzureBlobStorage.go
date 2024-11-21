package Request

// Details for loaction of file in Azure Blob Storage
type RequestAzureBlobStorage struct {

	// Blob Storage Container
	Container string

	// Azure File Name
	FileName string
}
