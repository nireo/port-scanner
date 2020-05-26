package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// CheckIfPortInString checks for a port number in a address and returns a boolean
func CheckIfPortInString(address string) bool {
	if strings.Index(address, ":") != -1 {
		return false
	}

	return true
}

// AddPortToAddress adds a port at the end of an address
func AddPortToAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

// SeparatePortsFromString returns ports from a string with ports separated by a , i.e 3000,8080,1234,3456
func SeparatePortsFromString(portString string) []int {
	splittedPortString := strings.Split(portString, ",")
	if len(splittedPortString) < 1 {
		fmt.Println("Please provide ports separated by a comma (,)")
		return []int{}
	}

	var portsArray []int
	for _, port := range splittedPortString {
		portStringToInt, err := strconv.Atoi(port)
		if err != nil {
			fmt.Printf("Problem parsing array port number:%s", port)
		}

		portsArray = append(portsArray, portStringToInt)
	}
	return portsArray
}

// ParseArgumentBoolean parses a string for a boolean value, then returns that value
func ParseArgumentBoolean(arg string) bool {
	if arg == "true" {
		return true
	}

	return false
}
