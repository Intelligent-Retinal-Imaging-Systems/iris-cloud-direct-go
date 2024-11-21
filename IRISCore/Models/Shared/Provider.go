package Shared

// Structure used to specify provider details
type Provider struct {
	// 10 digit National Provider Id
	NPI string
	// Providers name
	Name PersonName
	// Associations
	Associations string
	// Degrees
	Degrees string
	// Primary Taxonomy
	Taxonomy string
	// Email Address
	Email string
}
