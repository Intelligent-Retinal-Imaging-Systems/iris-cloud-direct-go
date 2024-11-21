package Shared

// Address represents a site's address.
type Address struct {
	Street string
	City   string
	Zip    string
}

// Site represents a request site with details.
type Site struct {
	LocalId string  // Site identifier
	Name    string  // Site name
	Address Address // Address details
}
