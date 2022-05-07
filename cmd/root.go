/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)
var (
	UsageLongDescription = "Maybe, we don't need the long description."
)
var rootCmd = &cobra.Command{
	Use:   "cli-x",
	Short: "😀 Cli-x is a cmd tool.",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😋 😍 🥰 You have installed cli-x successfully !")
	},
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(VersionCmd) 
	rootCmd.AddCommand(ServeCmd)
	rootCmd.AddCommand(ListCmd) 
	rootCmd.AddCommand(HttpCmd)
}


