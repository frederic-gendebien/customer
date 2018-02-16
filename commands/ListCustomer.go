package commands

import (
	"fmt"
	"github.com/frederic-gendebien/customer/persistence"
)

type ListCustomer struct{}

func NewListCustomer() *ListCustomer {
	return new(ListCustomer)
}

func (command *ListCustomer) Name() string {
	return "list"
}

func (command *ListCustomer) Description() string {
	return "list all customers aliases"
}

func (command *ListCustomer) Execute(arguments []string) error {
	aliases, err := persistence.List()
	if err != nil {
		return err
	}

	for _, alias := range aliases {
		fmt.Println(alias)
	}

	return nil
}
