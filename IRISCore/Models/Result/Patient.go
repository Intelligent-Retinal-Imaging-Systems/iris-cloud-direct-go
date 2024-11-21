package Result

import "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Shared"

// Structure containing Patient details
type Patient struct {
	// Id as specified by the client.  Typically MRN
	LocalId string
	// Patient Name details
	Name Shared.PersonName
	// Patient Date of birth
	Dob string
	// Patient Gender [Obsolete("Use Genders")]
	Gender int
	// Patient phone number
	Phone string
	// List of person genders
	Genders []Shared.PersonGender
	// IRIS-assigned PatientId
	PatientId int
}
