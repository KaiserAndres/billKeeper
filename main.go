package main

import (
	"os"
	"fmt"
	"errors"
	"github.com/urfave/cli"
)

var (
	notImplemented error = errors.New("This hasn't been implemented.")
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
	{
		Name: "list",
		Aliases: []string{"l"},
		Usage: "List all unpayed bills",
		Action: func(c *cli.Context) error {
			//TODO list all unpayed bills
			return notImplemented
		},
	},
	{
		Name: "add",
		Aliases: []string{"a"},
		Usage: "Add a new bill to the bill list"
		Action: func(c *cli.Context) error {
			return notIMplemented
		}
	},
	}
}
