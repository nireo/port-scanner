package main

import (
	"fmt"

	"github.com/nireo/port-scanner/cli"
	"github.com/nireo/port-scanner/scanner"
)

func main() {
	address, ports := cli.ParseCLIArguments()
	if address == "" {
		fmt.Println("ERROR: You need to provide an address to scan!")
		return
	}

	openPorts := scanner.RunScanner(address, ports)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
