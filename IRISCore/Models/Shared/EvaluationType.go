package Shared

type EvaluationType string

const (
	DiabeticRetinopathy       EvaluationType = "DR"
	Glaucoma                  EvaluationType = "Glaucoma"
	HIV                       EvaluationType = "HIV"
	MacularEdema              EvaluationType = "ME"
	AMD                       EvaluationType = "AMD"
	DiabeticRetinopathyAndAMD EvaluationType = "DR_AMD"
)
