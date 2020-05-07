package main // executable commands must always use package main.

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)

	ownerPtr := createCommand.String("owner", "Darth Vader", "Account owner. (Required)")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("a subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "create":
		createCommand.Parse(os.Args[2:])
		fmt.Println(" create account for :", *ownerPtr)
		fmt.Println("  tail:", createCommand.Args())

	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
