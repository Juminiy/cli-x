/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the usage of cli-x",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}
