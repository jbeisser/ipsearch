package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ipsearch"
)

func main() {
	// get IP from argv
	//ip := os.Args[1]
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	s = strings.TrimSuffix(s, "\n")
	ips := strings.Split(s, " ")

	for p := range ips {
		f := ipsearch.InternetDBLookup(ips[p])
		fmt.Printf("IP: %s\n", f.Ip)

		for p := range f.Ports {
			fmt.Println(f.Ports[p])
		}
	}
}
