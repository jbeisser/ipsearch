package main

import (
	"bufio"
	"fmt"
	"ipsearch"
	"os"
	"strings"
)

func main() {
	var in []string

	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		in = strings.Split(s, " ")

	} else {
		in = os.Args[1:]
	}

	for p := range in {
		f := ipsearch.InternetDBLookup(in[p])
		fmt.Printf("IP: %s\n", f.Ip)
		fmt.Printf("Hostnames: %s\n", f.Hostnames)

		for p := range f.Ports {
			fmt.Println(f.Ports[p])
		}
	}
}
