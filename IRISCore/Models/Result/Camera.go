package Result

// structure containing details of camera responsible for image
type Camera struct {
	// Camera Name - If not explicitely set, is serial number
	Name string
	// Id of Camera as specified by the submission organization
	LocalId string
	// Location of camera.  Used in matching process if other identifiers are not specified
	Location string
	// IPAddress of camera.  Used in matching process if other identifiers are not specified
	IPAddress string
	// MACAddress of camera NIC.
	MACAddress string
	// Serial number of camera
	SerialNumber string
	// Manufacturer of the device
	Manufacturer string
	// Model of device as specified by the manufacturer
	Model string
	// Software version of the device
	SoftwareVersion string
}
