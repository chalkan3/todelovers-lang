package commands

import (
	todelovers "mary_guica/pkg/todelovers_lang_tools.go"

	"github.com/spf13/cobra"
)

type Run struct {
	root *cobra.Command
}

func (r *Run) RegisterRunFunc() *cobra.Command {
	r.flags()
	r.root.Run = r.Run
	return r.root
}

func (r *Run) Run(cmd *cobra.Command, args []string) {
	bin, _ := cmd.Flags().GetString("bin")
	todelovers.New().Run(bin)
}

func (r *Run) flags() {
	r.root.Flags().String("bin", "todbin", "Name of bin")
}

func NewRun() Commands {
	return &Run{
		root: &cobra.Command{
			Use:   "run",
			Short: "Print the version number of mycli",
		},
	}
}
