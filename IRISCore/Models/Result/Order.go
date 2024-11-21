package Result

// Structure containing Order details for the results
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
	SingleEyeOnly bool
	// If the order is to have only one eye, this specifies the reason
	MissingEyeReason string
	// Unspecified additional information to include in the order
	AdditionalInfo string
	// If desired include the Encounter number
	EncounterNumber string
	// Specifies one or more evaluation types to process for this order
	EvaluationTypes []int
	/// Typically a value from DICOM
	StudyInstanceUniqueId string
	// Department Id associated with order
	DepartmentId string
	// Open field for anything
	OrderableIdentifier string
	// If true, order was expedited
	Expedite bool
	// Order Id as specified by IRIS
	PatientOrderId int
	// Date/Time service/exam was performed
	ServicedTime string
	// Status: AwaitingImages, QueuedForGrading, CheckedOutForGrading, Graded, OrderCancelled, UnderReview, Error, Complete
	Status string
}
