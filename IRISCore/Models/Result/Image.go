package Result

// Contains data about a single image
type Image struct {
	// Id of image as specified by the submitting organization.
	LocalId string
	// Timestamp when Image was taken
	Taken string
	// Laterality of image
	Laterality int
	// Context of Image
	ImageContext string
	// Optional Grouping context
	GroupId int
	// Optional Ordinal of image in group context by results receiving organization
	GroupOrdinal int
	// If specified this is the Id of the parent image for this image as specified by submitting image source
	ParentLocalId string
	// Id of the image in the IRIS database.
	OrderImageID int
	// Information about the Camera
	Camera Camera
	// Filename of image
	FileName string
	// Timestamp when image was received by Iris
	Received string
	// Container name that the file is stored in
	ContainerName string
}
