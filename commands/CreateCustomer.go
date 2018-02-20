package commands

import (
	"github.com/frederic-gendebien/customer/domain"
	"github.com/frederic-gendebien/customer/persistence"
)

type CreateCustomer struct {
	alias   string
	name    string
	address string
	vat     string
}

type CreateCustomerBuilder struct {
	Alias   string
	Name    string
	Address string
	Vat     string
}

func NewCreateCustomer(builder CreateCustomerBuilder) *CreateCustomer {
	command := new(CreateCustomer)
	command.alias 	= builder.Alias
	command.name 	= builder.Name
	command.address = builder.Address
	command.vat 	= builder.Vat

	return command
}

func (command *CreateCustomer) Execute() error {
	customer, err := command.toDomainObject()
	if err != nil {
		return err
	}

	return persistence.Save(command.alias, customer)
}

func (command *CreateCustomer) toDomainObject() (*domain.Customer, error) {
	return domain.NewCustomer(domain.CustomerBuilder{
		Name:    command.name,
		Address: command.address,
		Vat:     command.vat,
	})
}
