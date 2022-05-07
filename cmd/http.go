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
func testUri(uri string) (bool,error){ 
	match,err := regexp.MatchString(uri,
		`^((http:\/\/)|(https:\/\/))?([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}(\/)` )
	return match,err 
}
// httpCmd represents the http command
var HttpCmd = &cobra.Command{
	Use:   "http [URI] ",
	Short: "🚀 Cli-x http request",
	Long: UsageLongDescription,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		uri := args[0]
		if mtch,err := testUri(uri) ; err == nil && mtch { 
			fmt.Println("Http call to",uri)
			client := resty.New()
			resp,err := client.R().EnableTrace().Get(args[0]) 
			if err != nil {
				fmt.Println("Http request errors : ",err)
			} 
			fmt.Println(resp)
		} else { 
			if err != nil {
				fmt.Println("Error occurs ",err)
			}
			if !mtch  {
				fmt.Println("URI '", uri ,"' maybe not correct,please inspect input URI.")
			}
		}
		
	},
}
