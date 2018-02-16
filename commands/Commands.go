package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Command interface {
	Name() string
	Description() string
	Execute(arguments []string) error
}

var commands = []Command{
	NewCreateCustomer(),
	NewDeleteCustomer(),
	NewGetCustomer(),
	NewListCustomer(),
}

var commandsMap = make(map[string]Command)

func init() {
	arguments := os.Args
	if len(arguments) < 2 {
		menu(arguments[0])
		os.Exit(1)
	}

	for _, command := range commands {
		commandsMap[command.Name()] = command
	}
}

func Handle() error {
	arguments := os.Args
	commandName := arguments[1]
	command := commandsMap[commandName]
	if command == nil {
		return fmt.Errorf("invalid command: %s", commandName)
	}

	return command.Execute(arguments[2:])
}

func menu(programName string) {
	fmt.Println(programName, "<option> [option flags]")
	for _, command := range commands {
		fmt.Printf("\t%8s\t%s\n", command.Name(), command.Description())
	}
}

func ParseArguments(flagSet *flag.FlagSet, arguments []string) {
	err := flagSet.Parse(arguments)
	if err != nil {
		log.Fatal("could not parse: ", err)
	}
}

func AssertMandatory(flagSet *flag.FlagSet, arguments ...*string) {
	for _, argument := range arguments {
		if *argument == "" {
			flagSet.PrintDefaults()
			os.Exit(1)
		}
	}
}
