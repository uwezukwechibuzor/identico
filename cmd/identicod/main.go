package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/uwezukwechibuzor/Identico/app"
	"github.com/uwezukwechibuzor/Identico/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"identico",
		"cosmos",
		app.DefaultNodeHome,
		"identico",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
