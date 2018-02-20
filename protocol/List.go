package protocol

import (
	"github.com/frederic-gendebien/customer/commands"
)

type list struct {}

func NewList() *list {
	return new(list)
}

func (list *list) Name() string {
	return "list"
}

func (list *list) Description() string {
	return "list all customers aliases"
}

func (list *list) Parse(arguments []string) (commands.Command, error) {
	return commands.NewListCustomer(), nil
}
