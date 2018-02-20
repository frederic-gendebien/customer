package protocol

import (
	"flag"
	"fmt"

	"github.com/frederic-gendebien/customer/commands"
)

type create struct {}

func NewCreate() *create {
	return new(create)
}

func (create *create) Name() string {
	return "create"
}

func (create *create) Description() string {
	return "create a new customer"
}

func (create *create) Parse(arguments []string) (commands.Command, error) {
	flags 	:= flag.NewFlagSet("createCustomer", flag.ExitOnError)
	name 	:= flags.String("name", "", "name of the customer (mandatory)")
	address := flags.String("address", "", "address of the customer (mandatory)")
	vat 	:= flags.String("vat", "", "address of the customer (optional)")

	ParseArguments(flags, arguments)
	AssertMandatory(flags, name, address)

	options := flags.Args()
	if len(options) != 1 {
		return nil, fmt.Errorf("could not determine alias")
	}

	alias := options[0]
	if alias == "" {
		return nil, fmt.Errorf("could not work with empty alias")
	}

	command := commands.NewCreateCustomer(commands.CreateCustomerBuilder{
		Alias: 		alias,
		Name: 		*name,
		Address:	*address,
		Vat:		*vat,
	})

	return command, nil
}
