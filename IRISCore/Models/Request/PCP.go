package Request

import "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Shared"

type PCP struct {

	// Providers Name
	Name Shared.PersonName

	// Providers National ID number
	NPI string

	// Email Address for results
	EmailAddress string

	// Fax number for results
	FaxNumber string
}
