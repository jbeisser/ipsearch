package main

import (
	"fmt"
	"os"

	"ipsearch"
)

func main() {
	// get IP from stdin
	ip := os.Args[1]
	f := ipsearch.InternetDBLookup(ip)
	fmt.Println(f.Ip)
	for p := range f.Ports {
		fmt.Println(f.Ports[p])
	}

}
