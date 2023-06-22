package root

import (
	"fmt"
	"os"

	"github.com/mauricio-b-r/polygon-edge/command/backup"
	"github.com/mauricio-b-r/polygon-edge/command/genesis"
	"github.com/mauricio-b-r/polygon-edge/command/helper"
	"github.com/mauricio-b-r/polygon-edge/command/ibft"
	"github.com/mauricio-b-r/polygon-edge/command/license"
	"github.com/mauricio-b-r/polygon-edge/command/monitor"
	"github.com/mauricio-b-r/polygon-edge/command/peers"
	"github.com/mauricio-b-r/polygon-edge/command/secrets"
	"github.com/mauricio-b-r/polygon-edge/command/server"
	"github.com/mauricio-b-r/polygon-edge/command/status"
	"github.com/mauricio-b-r/polygon-edge/command/txpool"
	"github.com/mauricio-b-r/polygon-edge/command/version"
	"github.com/mauricio-b-r/polygon-edge/command/whitelist"
	"github.com/spf13/cobra"
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
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
