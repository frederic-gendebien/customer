package commands

import (
	"fmt"
	"github.com/frederic-gendebien/customer/persistence"
)

type ListCustomer struct {}

func NewListCustomer() *ListCustomer {
	return new(ListCustomer)
}

func (command *ListCustomer) Execute() error {
	aliases, err := persistence.List()
	if err != nil {
		return err
	}

	for _, alias := range aliases {
		fmt.Println(alias)
	}

	return nil
}
