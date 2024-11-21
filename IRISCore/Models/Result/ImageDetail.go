package Result

// /  Aggregate data for images associated with the order
type ImageDetails struct {
	// Total OS images
	LeftEyeCount int
	// Number of enhanced OD images
	LeftEyeEnhancedCount int
	// Number of original OS images
	LeftEyeOriginalCount int
	// Total OD images
	RightEyeCount int
	// Number of enhanced OD images
	RightEyeEnhancedCount int
	// Number of original OD images
	RightEyeOriginalCount int
	// If true the order only contains images from a single eye.
	SingleEyeOnly bool
	/// Total count of images associated with the order
	TotalCount int
}
