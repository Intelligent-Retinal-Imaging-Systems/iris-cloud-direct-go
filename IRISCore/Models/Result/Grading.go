package Result

import "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Shared"

// Structure containing grading results
type Grading struct {
	// Actual care plan for the patient
	CarePlanDescription string
	// Title of Care Plan (Should not be used to identify any level of pathology)
	CarePlanName string
	// Obsolete - See DiagnosisCodes
	DxCodes []string
	// Collection of codes associated with the order
	DiagnosisCodes []DiagnosisCode
	// Date/Time (UTC) when order was graded
	GradedTime string
	// Collection of notes (zero or more) created by the graded
	Notes []Note
	// Gradings for OD Laterality
	OD EyeSideGrading
	// Gradings for OS Laterality
	OS EyeSideGrading
	// If true the exam contains pathology findings
	Pathology bool
	// Grading Provider
	Provider Shared.Provider
	// One or more findings qualify as Urgent
	Urgent bool
}
