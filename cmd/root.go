/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ipsearch/shodan"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipsearch",
	Short: "Search details on an IP address",
	Long:  `Search Shodan's InternetDB database for details on an IP address.`,
	Run:   IpSearch,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("trace", "t", false, "trace result hostnames via GeoNET")
}

func IpSearch(cmd *cobra.Command, args []string) {
	traceFlag, err := cmd.Flags().GetBool("trace")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	f, err := shodan.InternetDBLookup(args[0])
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	if f.Ip != "" {
		fmt.Printf("IP: %s\n", f.Ip)

		if len(f.Hostnames) != 0 {
			fmt.Printf("Hostnames: %s\n", f.Hostnames)
			if traceFlag {
				for h := range f.Hostnames {
					hn, err := shodan.GeoDnsLookup(f.Hostnames[h])
					if err != nil {
						fmt.Println("error: %s", err)
					}
					for host := range *hn {
						fmt.Printf("Hostname: %s\n", f.Hostnames[h])
						fmt.Printf("City: %s\n", (*hn)[host].FromLoc.City)
						fmt.Printf("Country: %s\n", (*hn)[host].FromLoc.Country)
						fmt.Printf("LatLon: %s\n", (*hn)[host].FromLoc.Latlon)
						fmt.Println("Answers:")
						for a := range (*hn)[host].Answers {
							fmt.Printf("%s: %s\n", (*hn)[host].Answers[a].Type, (*hn)[host].Answers[a].Value)
						}
					}
				}
			}
		}

		if len(f.Ports) != 0 {
			fmt.Println("Ports:")
			for p := range f.Ports {
				fmt.Println(f.Ports[p])
			}
			fmt.Println("")
		}

		if len(f.Vulns) != 0 {
			fmt.Println("CVEs:")
			for v := range f.Vulns {
				fmt.Println(f.Vulns[v])
			}
			fmt.Println()
		}
	}
}
