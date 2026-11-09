/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "prints a hello world",
	Long: `Just call a println("Hello World!") man`,
	Run: func(cmd *cobra.Command, args []string) {
		exclamationsCount, _ := cmd.Flags().GetInt("exclamationsCount")
		name, _ := cmd.Flags().GetString("name")

		hello := fmt.Sprintf("Hello %s", name)

		for range exclamationsCount {
			hello = strings.Join([]string{hello, "!"}, "")
		}

		fmt.Printf("%s\n", hello)
	},
	PreRun: func(cmd *cobra.Command, args []string){
		fmt.Println("PreRun: Hello World!")
	},
	PostRun: func(cmd *cobra.Command, args []string){
		fmt.Println("PostRun: Hello World!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().IntP("exclamationsCount", "c", 1, "Escolha o numero de ! que tem no hello world")
	helloCmd.Flags().StringP("name", "n", "World", "Escolha o nome pra quem saudar")
}
