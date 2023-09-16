package runtime

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
	target, _ := cmd.Flags().GetString("target")
	todelovers.New().BuildAndRun(target)
}

func (r *Run) flags() {
	r.root.Flags().String("target", "main.todelovers", "Name of bin")
}

func NewRun() Commands {
	return &Run{
		root: &cobra.Command{
			Use:   "run",
			Short: "Print the version number of mycli",
		},
	}
}
