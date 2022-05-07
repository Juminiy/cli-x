/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "😛 Cli-x serves",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}
 
