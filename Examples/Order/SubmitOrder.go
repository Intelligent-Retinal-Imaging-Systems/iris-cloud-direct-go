package main

import (
	"iris-cloud-direct-go/IRISCore/Models/Config"
	"iris-cloud-direct-go/IRISCore/Models/Request"
	"iris-cloud-direct-go/IRISCore/Models/Shared"
	OrderSubmitter "iris-cloud-direct-go/IRISCore/Order"
	"log"
)

func main() {
	config, err := Config.LoadConfiguration("../configuration.json")
	if err != nil {
		log.Fatalf("Failed to load configuration from 'configuration.json': %v", err)
	}

	var order = Request.NewOrderRequest(config.ClientGuid, string(Shared.DiabeticRetinopathy), "NW")

	order.Site = Shared.Site{LocalId: "2134"}
	order.Order.LocalId = "399403"
	order.Patient.LocalId = "0001343"
	order.Patient.Name = Shared.PersonName{First: "Test", Last: "Patient"}
	order.Patient.Dob = "1/1/2000"
	order.Patient.Genders = []Shared.PersonGender{{Context: "BirthGender", Gender: "M"}}
	order.HealthPlan.LocalId = "123004"
	order.HealthPlan.Name = "EIeelPlan"
	order.HealthPlan.MemberId = "12304343"

	// Submit the order using the SubmitOrder method
	err = OrderSubmitter.SubmitOrder(config.ServiceBus.Orders, order)
	if err != nil {
		log.Fatalf("Failed to submit order: %v", err)
	} else {
		log.Println("Order successfully submitted")
	}
}
