/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/spf13/cobra"
	"github.com/go-resty/resty/v2"
)
// regex
func testUri(uri string) (bool,string){ 
	if ! ( strings.HasPrefix(uri,"http://") || strings.HasPrefix(uri,"https://") ) { 
		uri = strings.Join([]string{"https://",uri},"") 
	}
	validUri,_ := regexp.Compile(
		`^((http:\/\/)|(https:\/\/))?([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}`) 
	return validUri.MatchString(uri),uri 


}
func httpGet(cmd *cobra.Command, args []string) { 
	uri := args[0] 
	if mtch,uri := testUri(uri) ; mtch {
		fmt.Println("Http GET call to",uri) 
		client := resty.New() 
		resp,err := client.R().EnableTrace().Get(uri) 
		if err != nil {
			fmt.Println("Http request errors : ",err)
		} 
		fmt.Println(resp)
	} else {
		fmt.Println("URI '", uri ,"' maybe not correct,please inspect input URI.")
	} 
} 

func httpPost(cmd *cobra.Command, args []string) {
	uri := args[0] 
	if mtch,uri := testUri(uri) ; mtch {
		fmt.Println("Http POST call to",uri) 
		client := resty.New() 
		resp,err := client.R().
			SetHeader("Content-Type","application/json").
			SetBody(args[1]).
			Post(uri) 
		if err != nil { 
			fmt.Println("Http request errors : ",err)
		} 
		fmt.Println(resp)
	} else {
		fmt.Println("URI '", uri ,"' maybe not correct,please inspect input URI.")
	} 
}
// httpCmd represents the http command
var HttpCmd = &cobra.Command{
	Use:   "http [uri]",
	Short: "🚀 Cli-x http request",
	Long: UsageLongDescription,
	Args: cobra.MinimumNArgs(1),
	Run: httpGet ,
}
var getCmd = &cobra.Command {
	Use: "get [uri]",
	Short: "🚀 🚀 ✈ Cli-x http get request",
	Long: UsageLongDescription ,
	Args: cobra.MinimumNArgs(1) ,
	Run: httpGet , 
}
var postCmd = &cobra.Command {
	Use: "post [uri] [# body]",
	Short: "🚀 🚀 🛰 Cli-x http post request",
	Long: UsageLongDescription ,
	Args: cobra.MinimumNArgs(2) ,
	Run: httpPost , 
}