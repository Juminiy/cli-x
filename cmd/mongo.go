/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "chisato.live/cli-x/connect"
)

var MongoCmd = &cobra.Command{
	Use:   "mongo",
	Short: "😏 To store data in mongodb server.",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mongo called")
	},
}
