package cli

import (
	"flag"
)

// ParseCLIArguments parses the CLI arguments and returns them
func ParseCLIArguments() (string, string) {
	var portsString string
	flag.StringVar(&portsString, "ports", "", "enter specific ports to be scanned separated by commas i.e: 1,8008,3001")

	var addressFlag string
	flag.StringVar(&addressFlag, "address", "", "enter the address you want to scan")

	return addressFlag, portsString
}
