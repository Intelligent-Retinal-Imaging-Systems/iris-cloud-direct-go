package Result

// Iris defined site associated with the order
type ResultSite struct {
	// Identifier as specified by IRIS
	ClientId int
	// Site name
	ClientName string
	// Timezone site prefers
	ClientTimeZone string
	// City where site is located
	City string
	// US State where site is located (can determine grading assignments)
	State string
	// Zip code where site is located
	Zip string
	// Country where site is located
	Country string
	// Unique identifier of site as assigned by IRIS
	PracticeGuid string
	// Area code portion of phone number for site
	AreaCode string
	// Phone number of site
	PhoneNumber string
	// Id of Organization Site operates under as defined by IRIS
	OrganizationId int
	// Name of organiztion site operates under
	OrganizationName string
}
