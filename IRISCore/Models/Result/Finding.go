package Result

// Structure containing an individual finding.  A Finding can either be Qualified or stand alone.
// For example: Suspected Glaucoma is not Qualified thus Result is empty
// Diabetic Retinopathy will always be qualified as it is a diagnosis, not just a suspected condition
type Finding struct {
	// Finding / Disease evaluated
	Finding string
	// Result / Qualifier  (e.g.: None, Mild, Positive, etc...)
	Result string
}
