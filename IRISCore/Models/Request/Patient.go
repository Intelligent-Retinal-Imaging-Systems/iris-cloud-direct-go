package Request

import "iris-cloud-direct-go/IRISCore/Models/Shared"

// Patient details for order
type Patient struct {

	// Id as specified by the client.  Typically MRN
	LocalId string

	// Patient Name details
	Name Shared.PersonName

	// Patient Date of birth
	Dob string

	// Patient Gender -- Obsolete("Use Genders")
	Gender string

	// Patient phone number
	Phone string

	// List of person genders
	Genders []Shared.PersonGender

	// Gender Race code
	// 1002-5: American Indian or Alaska Native
	// 2028-9: Asian
	// 2054-5: Black or African American
	// 2076-8: Native Hawaiian or Other Pacific Islander
	// 2131-1: Other Race
	// 2106-3: White
	Race string

	// Ethnicity code
	Ethnicity string

	// Language Code es=Spanish, fr=French, en-US=English(United States), etc...
	PrimaryLanguage string

	// Marital Status single letter code
	MaritalStatus string

	// Email address for patient
	Email string

	// Free form additonal information specific to the patient
	AdditionalInfo string

	// Patient Address Details
	Address Shared.Address

	// ICD-10 Code diagnosed that supports the exam
	DxCode string
}
