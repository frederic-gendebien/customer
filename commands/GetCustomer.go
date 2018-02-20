package commands

import (
	"fmt"
	"github.com/frederic-gendebien/customer/persistence"
	"github.com/frederic-gendebien/customer/serialization"
)

type GetCustomer struct {
	alias string
	format string
}

type GetCustomerBuilder struct {
	Alias string
	Format string
}

func NewGetCustomer(builder GetCustomerBuilder) *GetCustomer {
	command := new(GetCustomer)
	command.alias = builder.Alias
	command.format = builder.Format

	return command
}

func (command *GetCustomer) Execute() error {
	customer, err := persistence.Get(command.alias)
	if err != nil {
		return err
	}

	var serializer serialization.Serializer
	switch command.format {
	case "protobuf"	: serializer = serialization.Protobuf.Serializer(); break
	case "yaml"		: serializer = serialization.Yaml.Serializer(); break
	case "json"		: serializer = serialization.Json.Serializer(); break
	default 		: serializer = serialization.Json.Serializer(); break
	}

	bytes, err := serializer.Marshal(customer)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
