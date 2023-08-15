package visitor

import "fmt"

type Customer interface {
	Accpet(visitor Visitor)
}

type Visitor interface {
	Visit(customer Customer)
}

type EnterpriseCustomer struct {
	Name string
}

func (customer *EnterpriseCustomer) Accpet(visitor Visitor) {
	visitor.Visit(customer)
}

type IndividualCustomer struct {
	Name string
}

func (customer *IndividualCustomer) Accpet(visitor Visitor) {
	visitor.Visit(customer)
}

type ServiceRequestVisitor struct {
}

func (visitor *ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.Name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.Name)
	}
}

type AnalysisVisitor struct {
}

func (visitor *AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.Name)
	}
}
