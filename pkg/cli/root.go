package cli

import (
	"fmt"
	"mary_guica/pkg/cli/internal/commands"
	"os"

	"github.com/spf13/cobra"
)

type TodeLoverCLI struct {
	rootCMD *cobra.Command
}

func (n *TodeLoverCLI) registerCommands() {
	for _, cmd := range []commands.Commands{
		commands.NewBuild(),
		commands.NewRun(),
	} {
		n.rootCMD.AddCommand(cmd.RegisterRunFunc())
	}
}

func (n *TodeLoverCLI) Run() {
	n.registerCommands()
	if err := n.rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func NewTodeLoverCLI() *TodeLoverCLI {
	root := &cobra.Command{Use: "tode-lovers"}
	return &TodeLoverCLI{
		rootCMD: root,
	}
}
