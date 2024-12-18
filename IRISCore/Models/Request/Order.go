package Request

// Contains Details for the order
type Order struct {

	// Date/Time offset when order was created
	CreatedTime string

	// Date/Time offset when order is scheduled for image capture
	ScheduledTime string

	// Id for the order as specified by the submission source
	LocalId string

	// US State where the images are being taken.  Can be full name or 2 character abbreviation
	State string

	// If true images will be taken for a single eye only
	// Setting this value to true will prevent an order from delaying in wait of images for both eyes
	SingleEyeOnly bool

	// If the order is to have only one eye, this specifies the reason
	MissingEyeReason string

	// Unspecified additional information to include in the order
	AdditionalInfo string

	// If desired include the Encounter number
	EncounterNumber string

	// Specifies one or more evaluation types of the following Options
	// DR=Diabetic Retinopathy exam, Glaucoma=Glaucoma exam, HIV=HIV Retinopathy exam, AMD=Age-Related Macular Degeneration exam,
	// DR_AMD=Combined Diabetic Retinopathy & Age-Related Macular Degeneration exam
	EvaluationTypes []string

	// Typically a value from DICOM
	StudyInstanceUniqueId string

	// Department Id associated with order
	DepartmentId string

	// Open field for specific identifier of order. This is for support of EMR/EHR systems that use this concept
	OrderableIdentifier string
	// Procedure code for order

	CPTCode string
}
