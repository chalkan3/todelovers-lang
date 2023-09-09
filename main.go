package main

import (
	"fmt"
	"io/ioutil"
	"mary_guica/engine"
	"mary_guica/tvm"
	"os"

	"github.com/spf13/cobra"
)

// Arrumar a arvore sintaxica
// Realizar a criação do codigo intermediario
// fazer a  geração de código objeto com assembly

func main() {
	var rootCmd = &cobra.Command{Use: "tode-lovers"}

	var builCmd = &cobra.Command{
		Use:   "build",
		Short: "Print the version number of mycli",
		Run: func(cmd *cobra.Command, args []string) {
			target, _ := cmd.Flags().GetString("target")

			dsl, err := engine.File(target)
			if err != nil {
				panic(err)
			}

			symbleTable := engine.NewSymbolTable()
			lexer := engine.NewLexer(dsl).Tokenize()
			nodeFactory := engine.NewNodeFactory()
			assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(false)

			root := assembler.GetRoot()
			root.RegisterSymbols(symbleTable, nil)

			code := root.GenerateIntermediateCode()
			// Sample data as a byte slice

			// Define the file path where you want to save the data
			filePath := "lovers.todbin"

			// Save the []byte to a file
			err = ioutil.WriteFile(filePath, code, 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Print the version number of mycli",
		Run: func(cmd *cobra.Command, args []string) {
			bin, _ := cmd.Flags().GetString("bin")

			dataRead, err := ioutil.ReadFile(bin)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			vm := tvm.NewTVM()

			tvm.LoadCode(vm, dataRead)

		},
	}

	builCmd.Flags().String("target", "main.todelovers", "Name of file")
	runCmd.Flags().String("bin", "main.todebin", "Name of bin")

	rootCmd.AddCommand(builCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
