runner

import "github.com/spf13/cobra"

type Commands interface {
	RegisterRunFunc() *cobra.Command
	Run(cmd *cobra.Command, args []string)
}
