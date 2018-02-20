package domain

import "fmt"

type Customer struct {
	name    string
	address string
	vat     string
}

type CustomerBuilder struct {
	Name    string
	Address string
	Vat     string
}

func NewCustomer(builder CustomerBuilder) (*Customer, error) {
	if builder.Name == "" {
		return nil, fmt.Errorf("could not create Customer without name")
	}
	if builder.Address == "" {
		return nil, fmt.Errorf("could not create Customer without address")
	}

	customer := new(Customer)
	customer.name 		= builder.Name
	customer.address 	= builder.Address
	customer.vat 		= builder.Vat

	return customer, nil
}

func (customer *Customer) GetName() string {
	return customer.name
}

func (customer *Customer) GetAddress() string {
	return customer.address
}

func (customer *Customer) GetVat() string {
	return customer.vat
}
