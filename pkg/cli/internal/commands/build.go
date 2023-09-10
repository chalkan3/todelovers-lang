package commands

import (
	todelovers "mary_guica/pkg/todelovers_lang_tools.go"

	"github.com/spf13/cobra"
)

type Build struct {
	root *cobra.Command
}

func (r *Build) RegisterRunFunc() *cobra.Command {
	r.root.Run = r.Run
	return r.root
}

func (r *Build) Run(cmd *cobra.Command, args []string) {
	r.flags()
	target, _ := cmd.Flags().GetString("target")
	todelovers.New().Build(target)
}

func (r *Build) flags() {
	r.root.Flags().String("target", "main.todelovers", "Name of file")
}

func NewBuild() Commands {
	return &Build{
		root: &cobra.Command{
			Use:   "build",
			Short: "Print the version number of mycli",
		},
	}
}
