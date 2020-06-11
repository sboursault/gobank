package main // executable commands must always use package main.

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

					if c.Args().Len() != 1 {
						cli.ShowCommandHelpAndExit(c, "open-account", 1)
					}

					owner := c.Args().Get(0)
					accountNo := bank.OpenAccount(owner)

					fmt.Printf("new account number: %+v\n", accountNo)

					return nil
				}},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposite",
				ArgsUsage: "ACCOUNT-NUMBER AMOUNT",
				Action: func(c *cli.Context) error {

					if c.Args().Len() != 2 {
						cli.ShowCommandHelpAndExit(c, "deposit", 1)
					}

					accountNumber := c.Args().Get(0)
					amount := strToFloat32(c.Args().Get(1))

					bank.Deposit(accountNumber, amount)

					return nil
				}},
			{
				Name:      "withdraw",
				Aliases:   []string{"w"},
				Usage:     "Make a withdrawal",
				ArgsUsage: "ACCOUNT-NUMBER AMOUNT",
				Action: func(c *cli.Context) error {

					if c.Args().Len() != 2 {
						cli.ShowCommandHelpAndExit(c, "withdraw", 1)
					}

					accountNumber := c.Args().Get(0)
					amount := strToFloat32(c.Args().Get(1))

					err := bank.Withdraw(accountNumber, amount)

					if err != nil {
						fmt.Println(err)
					}

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func strToFloat32(input string) float32 {
	f, err := strconv.ParseFloat(input, 32)
	checkErr(err)
	return float32(f)
}
