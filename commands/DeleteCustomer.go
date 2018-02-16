package commands

import (
	"fmt"
	"github.com/frederic-gendebien/customer/persistence"
)

type DeleteCustomer struct{}

type CustomerDeletion struct {
	alias string
}

func NewDeleteCustomer() *DeleteCustomer {
	return new(DeleteCustomer)
}

func (command *DeleteCustomer) Name() string {
	return "delete"
}

func (command *DeleteCustomer) Description() string {
	return "delete a user from his alias"
}

func (command *DeleteCustomer) Execute(arguments []string) error {
	customerDeletion, err := command.parse(arguments)
	if err != nil {
		return err
	}

	return customerDeletion.execute()
}

func (command *DeleteCustomer) parse(arguments []string) (*CustomerDeletion, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("wrong number of parameters")
	}

	alias := arguments[0]
	if alias == "" {
		return nil, fmt.Errorf("could not work with empty alias")
	}

	customerDeletion := CustomerDeletion{
		alias: alias,
	}

	return &customerDeletion, nil
}

func (customerDeletion *CustomerDeletion) execute() error {
	return persistence.Delete(customerDeletion.alias)
}
