package Shared

// Health plan details
type HealthPlan struct {
	// Id of HealthPlan as specified by the submitting organization
	LocalId string
	// Name of HealthPlan as specified by the submitting organization
	Name string
	// MemberId of patient as specified by the Healthplan
	MemberId string
	// Name of the Payer organization providing the health plan to the patient
	HealthPlanPayerName string
	// Group name of the patients health plan
	GroupName string
	// Group number of the patients health plan
	GroupNumber string
	// Effective date of the patients health plan
	EffectiveDate string
	// Expiration date of the patients health plan
	ExpirationDate string
	// Patients health plan policy number
	PolicyNumber string
}
