package protocol

import (
	"fmt"
	"flag"

	"github.com/frederic-gendebien/customer/commands"
)

type get struct{}

func NewGet() *get {
	return new(get)
}

func (get *get) Name() string {
	return "get"
}

func (get *get) Description() string {
	return "get a user from his alias"
}

func (get *get) Parse(arguments []string) (commands.Command, error) {
	flags := flag.NewFlagSet("getCustomer", flag.ExitOnError)
	format := flags.String("format", "json", "format to use (protobuf | json | yaml)")

	ParseArguments(flags, arguments)

	options := flags.Args()
	if len(options) != 1 {
		return nil, fmt.Errorf("wrong number of parameters")
	}

	alias := options[0]
	if alias == "" {
		return nil, fmt.Errorf("could not work with empty alias")
	}

	command := commands.NewGetCustomer(commands.GetCustomerBuilder{
		Alias: 	alias,
		Format: *format,
	})

	return command, nil
}
