/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"regexp"
	"github.com/spf13/cobra"
	"github.com/go-resty/resty/v2"
)
// regex
func testUri(uri string) (bool){ 
	validUri,_ := regexp.Compile(
		`^((http:\/\/)|(https:\/\/))?([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}`)
	return validUri.MatchString(uri) 
}
// httpCmd represents the http command
var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "🚀 Cli-x http request",
	Long: UsageLongDescription,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		uri := args[0] 
		if mtch := testUri(uri) ; mtch {
			fmt.Println("Http call to",uri)
			client := resty.New() 
			resp,err := client.R().EnableTrace().Get(args[0]) 
			if err != nil {
				fmt.Println("Http request errors : ",err)
			} 
			fmt.Println(resp)
		} else {
			fmt.Println("URI '", uri ,"' maybe not correct,please inspect input URI.")
		}
	},
}
