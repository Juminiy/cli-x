/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-x",
	Short: "CHISATO CLI-X",
	Long: `Chisato Cli-x`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-x")
	},
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(VersionCmd) 
	rootCmd.AddCommand(ServeCmd)
}


