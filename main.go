package main // executable commands must always use package main.

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/sboursault/gobank/bank"
)

func main() {
	app := &cli.App{
		Name:  "gobank",
		Usage: "A simplistic bank account service based on event sourcing",
		Commands: []*cli.Command{
			{
				Name:      "open-account",
				Aliases:   []string{"oa"},
				Usage:     "Opens a bank account",
				ArgsUsage: "OWNER",
				Action: func(c *cli.Context) error {

					owner := c.Args().Get(0)

					if owner == "" {
						cli.ShowCommandHelpAndExit(c, "open-account", 1)
					}

					accountNo := bank.OpenAccount(owner)

					fmt.Printf("new account number: %+v\n", accountNo)
					return nil
				},
			},
			{
				Name:    "deposit",
				Aliases: []string{"d"},
				Usage:   "Make a deposite",
				Action: func(c *cli.Context) error {

					accountNumber := c.Args().Get(0)
					amount := c.Args().Get(1)

					if accountNumber == "" || amount == "" {
						cli.ShowCommandHelpAndExit(c, "deposit", 1)
					}

					accountNo := bank.Deposit(accountNumber)

					fmt.Printf("new account number: %+v\n", accountNo)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
