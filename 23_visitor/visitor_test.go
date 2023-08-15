package visitor

import "testing"

func TestVisitor(t *testing.T) {
	var customers []Customer
	customers = append(customers, &EnterpriseCustomer{Name: "Alice"})
	customers = append(customers, &IndividualCustomer{Name: "Bob"})
	visitor1 := ServiceRequestVisitor{}
	visitor2 := AnalysisVisitor{}
	for i := 0; i < len(customers); i++ {
		customers[i].Accpet(&visitor1)
		customers[i].Accpet(&visitor2)
	}
}
