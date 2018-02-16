package commands

import (
	"flag"
	"fmt"
	"github.com/frederic-gendebien/customer/domain"
	"github.com/frederic-gendebien/customer/persistence"
)

type CreateCustomer struct{}

type CustomerCreation struct {
	alias   string
	name    string
	address string
	vat     string
}

func NewCreateCustomer() *CreateCustomer {
	return new(CreateCustomer)
}

func (command *CreateCustomer) Name() string {
	return "create"
}

func (command *CreateCustomer) Description() string {
	return "create a new customer"
}

func (command *CreateCustomer) Execute(arguments []string) error {
	customerCreation, err := command.parse(arguments)
	if err != nil {
		return err
	}

	return customerCreation.execute()
}

func (command *CreateCustomer) parse(arguments []string) (*CustomerCreation, error) {
	flags := flag.NewFlagSet("createCustomer", flag.ExitOnError)
	name := flags.String("name", "", "name of the customer (mandatory)")
	address := flags.String("address", "", "address of the customer (mandatory)")
	vat := flags.String("vat", "", "address of the customer (optional)")

	ParseArguments(flags, arguments)
	AssertMandatory(flags, name, address)

	options := flags.Args()
	if len(options) != 1 {
		return nil, fmt.Errorf("could not determine alias")
	}

	alias := &options[0]

	customerCreation := CustomerCreation{
		alias:   *alias,
		name:    *name,
		address: *address,
		vat:     *vat,
	}

	return &customerCreation, nil
}

func (command *CustomerCreation) execute() error {
	customer, err := command.toDomainObject()
	if err != nil {
		return err
	}

	return persistence.Save(command.alias, customer)
}

func (command *CustomerCreation) toDomainObject() (*domain.Customer, error) {
	return domain.NewCustomer(domain.CustomerBuilder{
		Name:    command.name,
		Address: command.address,
		Vat:     command.vat,
	})
}
