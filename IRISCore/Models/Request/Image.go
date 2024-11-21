package Request

type Image struct {
	// Id of image as specified by the submitting organization.
	LocalId string

	// Timestamp when Image was taken
	DateTimeOffset string

	// Laterality of image
	Laterality string

	// Context of Image
	ImageContext string

	// Optional Grouping context
	GroupId int

	// Optional Ordinal of image in group context by results receiving organization
	GroupOrdinal int

	// If specified this is the Id of the parent image for this image
	ParentLocalId string

	// Details of image location uploaded in Azure blob storage
	AzureBlobStorage RequestAzureBlobStorage
}

type ImageDictionary struct {
	Images []Image
}
