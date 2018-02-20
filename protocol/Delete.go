package protocol

import (
	"fmt"

	"github.com/frederic-gendebien/customer/commands"
)

type delete struct {}

func NewDelete() *delete {
	return new(delete)
}

func (delete *delete) Name() string {
	return "delete"
}

func (delete *delete) Description() string {
	return "delete a user from his alias"
}

func (delete *delete) Parse(arguments []string) (commands.Command, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("wrong number of parameters")
	}

	alias := arguments[0]
	if alias == "" {
		return nil, fmt.Errorf("could not work with empty alias")
	}

	command := commands.NewDeleteCustomer(commands.DeleteCustomerBuilder{
		Alias: alias,
	})

	return command, nil
}
