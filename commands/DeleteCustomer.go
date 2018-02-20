package commands

import (
	"github.com/frederic-gendebien/customer/persistence"
)

type DeleteCustomer struct {
	alias string
}

type DeleteCustomerBuilder struct {
	Alias string
}

func NewDeleteCustomer(builder DeleteCustomerBuilder) *DeleteCustomer {
	command := new(DeleteCustomer)
	command.alias = builder.Alias

	return command
}

func (command *DeleteCustomer) Execute() error {
	return persistence.Delete(command.alias)
}
