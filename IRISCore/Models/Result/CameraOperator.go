package Result

import "iris-cloud-direct-go/IRISCore/Models/Shared"

// Structure containing information about the Camera operator that captured the images.
type CameraOperator struct {
	// Camera operators name
	Name Shared.PersonName
	// Notes created by this camera operator relative to this document/order
	Notes []Note
	// User name as supplied on authentication to IRIS systems (email addr)
	UserName string
}
