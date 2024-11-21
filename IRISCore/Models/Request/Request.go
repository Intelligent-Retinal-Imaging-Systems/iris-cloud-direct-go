package Request

import "github.com/Intelligent-Retinal-Imaging-Systems/iris-cloud-direct-go/IRISCore/Models/Shared"

// Structure used to send an order to IRIS
type OrderRequest struct {

	// IRIS Request version
	Version string

	// Identifies the submission source organization
	ClientGuid string

	// Site the exam is being taken from
	Site Shared.Site

	// Username of the user submitting this order
	UserNameSubmitting string

	// Type of order request
	// NW = New Order, XO = Change order, CA = Cancel Order, ResendResult = Requesting resend of results
	OrderControlCode string

	// Camera associated with the request
	Camera Camera

	// Order specific data
	Order Order

	// Patient specific data
	Patient Patient

	// Medical provider ordering the exam
	OrderingProvider Shared.Provider

	// Medical provider referring patient for exam
	ReferringProvider Shared.Provider

	// Camera Operator performing the exam
	CameraOperator Shared.Provider

	// Username (login) for camera operator
	CameraOperatorUserName string

	// HealthPlan information associated with the request
	HealthPlan Shared.HealthPlan
}

// Return a new OrderRequest object with required fields
func NewOrderRequest(clientGuid string, evaluationType string, orderControlCode string) OrderRequest {
	return OrderRequest{
		Version:                "V2.3.1",
		ClientGuid:             clientGuid,
		Site:                   Shared.Site{},
		UserNameSubmitting:     "",
		OrderControlCode:       orderControlCode,
		Camera:                 Camera{},
		Order:                  Order{EvaluationTypes: []string{evaluationType}},
		Patient:                Patient{},
		OrderingProvider:       Shared.Provider{},
		ReferringProvider:      Shared.Provider{},
		CameraOperator:         Shared.Provider{},
		CameraOperatorUserName: "",
		HealthPlan:             Shared.HealthPlan{},
	}
}
