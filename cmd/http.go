/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/go-resty/resty/v2"

)

// httpCmd represents the http command
var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "🚀 Cli-x http request",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
		client := resty.New()
		resp,err := client.R().EnableTrace().Get("https://meowfacts.herokuapp.com") 
		if err != nil {
			fmt.Println("Http request errors : ",err)
		} 
		fmt.Println(resp)
	},
}
