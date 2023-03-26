package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	app "github.com/0xknstntn/quadrate/app"
	"github.com/0xknstntn/quadrate/cmd/quadrated/cmd"
	"github.com/0xknstntn/quadrate/types"
)

func main() {
	// Set address prefix and cointype
	types.SetConfig()

	err := sdk.RegisterDenom(types.DefaultDenom, sdk.NewDecWithPrec(1, types.DefaultDenomPrecision))
	if err != nil {
		panic(err)
	}

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
