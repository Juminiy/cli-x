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
	Short: "CHISATO CLI-X SERVE",
	Long: `Chisato Cli-x Serve`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}
 
