package root

import (
	"fmt"
	"os"

	"github.com/gitshock-labs/go-client/command/backup"
	"github.com/gitshock-labs/go-client/command/genesis"
	"github.com/gitshock-labs/go-client/command/helper"
	"github.com/gitshock-labs/go-client/command/ibft"
	"github.com/gitshock-labs/go-client/command/license"
	"github.com/gitshock-labs/go-client/command/loadbot"
	"github.com/gitshock-labs/go-client/command/monitor"
	"github.com/gitshock-labs/go-client/command/peers"
	"github.com/gitshock-labs/go-client/command/secrets"
	"github.com/gitshock-labs/go-client/command/server"
	"github.com/gitshock-labs/go-client/command/status"
	"github.com/gitshock-labs/go-client/command/txpool"
	"github.com/gitshock-labs/go-client/command/version"
	"github.com/gitshock-labs/go-client/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Gitshock Edgeware is a framework for building Ethereum-compatible Blockchain networks",
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
		loadbot.GetCommand(),
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
