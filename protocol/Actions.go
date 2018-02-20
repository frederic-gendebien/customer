package protocol

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/frederic-gendebien/customer/commands"
)

type Action interface {
	Name() string
	Description() string
	Parse(arguments []string) (commands.Command, error)
}

var (
	actions = []Action{
		NewCreate(),
		NewDelete(),
		NewGet(),
		NewList(),
	}

	actionMap = make(map[string]Action)
)

func init() {
	arguments := os.Args
	if len(arguments) < 2 {
		menu(arguments[0])
		os.Exit(1)
	}

	for _, command := range actions {
		actionMap[command.Name()] = command
	}
}

func Parse() (commands.Command, error) {
	arguments 	:= os.Args
	commandName := arguments[1]
	command 	:= actionMap[commandName]

	if command == nil {
		return nil, fmt.Errorf("invalid command: %s", commandName)
	}

	return command.Parse(arguments[2:])
}

func menu(programName string) {
	fmt.Println(programName, "<option> [option flags]")
	for _, command := range actions {
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
