package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Pawswap/posa/command/backup"
	"github.com/Pawswap/posa/command/bridge"
	"github.com/Pawswap/posa/command/genesis"
	"github.com/Pawswap/posa/command/helper"
	"github.com/Pawswap/posa/command/ibft"
	"github.com/Pawswap/posa/command/license"
	"github.com/Pawswap/posa/command/monitor"
	"github.com/Pawswap/posa/command/peers"
	"github.com/Pawswap/posa/command/polybft"
	"github.com/Pawswap/posa/command/polybftsecrets"
	"github.com/Pawswap/posa/command/regenesis"
	"github.com/Pawswap/posa/command/rootchain"
	"github.com/Pawswap/posa/command/secrets"
	"github.com/Pawswap/posa/command/server"
	"github.com/Pawswap/posa/command/status"
	"github.com/Pawswap/posa/command/txpool"
	"github.com/Pawswap/posa/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
