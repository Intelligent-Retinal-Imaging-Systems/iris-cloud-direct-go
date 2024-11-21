package Result

import "iris-cloud-direct-go/IRISCore/Models/Shared"

// Structure containing results of a graded order
type OrderResult struct {
	// Version of IRIS result structure
	Version string `json:"version"`
	// A GUID  uniquely identifying this document
	TransactionId string `json:"transactionId"`
	// Date/Time document was created
	Timestamp string `json:"timestamp"`
	// If there was an error generating this document, this contains details
	ErrorMessage string `json:"errorMessage"`
	// Specifies the type of result object
	ResultObjectType string
	// Operator performing the exam
	CameraOperator CameraOperator `json:"cameraOperator"`
	// Gradings captured for this order
	Gradings Grading `json:"gradings"`
	// Health plan associated with the results
	HealthPlan Shared.HealthPlan `json:"healthPlan"`
	// Summary data for images submitted to the order
	ImageDetails ImageDetails `json:"imageDetails"`
	// Images associated with the order
	Images []Image `json:"images"`
	// Details of the order
	Order Order `json:"order"`
	// Ordering Provider
	OrderingProvider Shared.Provider `json:"orderingProvider"`
	// Referring Provider
	ReferringProvider Shared.Provider `json:"referringProvider"`
	// Patient details
	Patient Patient `json:"patient"`
	// P=Preliminary, A=Addendum, F=Final, C=Correction, R=Resend
	ResultCode string `json:"resultCode"`
	// Contains the result document payload
	ResultsDocument Document `json:"resultsDocument"`
	// Location the Order is assigned to.
	Site ResultSite `json:"site"`
}
