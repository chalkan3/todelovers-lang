package commands

import (
	todelovers "mary_guica/pkg/todelovers_lang_tools.go"

	"github.com/spf13/cobra"
)

type Start struct {
	root *cobra.Command
}

func (r *Start) RegisterRunFunc() *cobra.Command {
	r.flags()
	r.root.Run = r.Run
	return r.root
}

func (r *Start) Run(cmd *cobra.Command, args []string) {
	bin, _ := cmd.Flags().GetString("bin")
	todelovers.New().Run(bin)
}

func (r *Start) flags() {
	r.root.Flags().String("bin", "todbin", "Name of bin")
}

func NewStart() Commands {
	return &Start{
		root: &cobra.Command{
			Use:   "start",
			Short: "Print the version number of mycli",
		},
	}
}
