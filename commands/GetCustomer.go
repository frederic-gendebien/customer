package commands

import (
	"fmt"
	"github.com/frederic-gendebien/customer/persistence"
	"github.com/frederic-gendebien/customer/serialization"
	"flag"
)

type GetCustomer struct{}

type CustomerGet struct {
	alias string
	format string
}

func NewGetCustomer() *GetCustomer {
	return new(GetCustomer)
}

func (command *GetCustomer) Name() string {
	return "get"
}

func (command *GetCustomer) Description() string {
	return "get a user from his alias"
}

func (command *GetCustomer) Execute(arguments []string) error {
	customerGet, err := command.parse(arguments)
	if err != nil {
		return err
	}

	return customerGet.execute()
}

func (command *GetCustomer) parse(arguments []string) (*CustomerGet, error) {
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

	customerGet := CustomerGet{
		alias: 	alias,
		format: *format,
	}

	return &customerGet, nil
}

func (command *CustomerGet) execute() error {
	customer, err := persistence.Get(command.alias)
	if err != nil {
		return err
	}

	var serializer serialization.Serializer
	switch command.format {
	case "protobuf"	: serializer = serialization.Protobuf.Serializer(); break
	case "json"		: serializer = serialization.Json.Serializer(); break
	case "yaml"		: serializer = serialization.Yaml.Serializer(); break
	}

	bytes, err := serializer.Marshal(customer)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
