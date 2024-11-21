package Result

// Structure containing grading results for one eye (OD / OS)
type EyeSideGrading struct {
	// All findings for this eye (if any)
	Findings []Finding
	// If false, the eye was not gradable
	Gradable bool
	// If ungradable this is a collection of reasons why it wasn't
	UngradableReasons []string
	// If no images were capture for this eye and a reason was given/determined, this will be set with that reason
	MissingEyeReason string
}
