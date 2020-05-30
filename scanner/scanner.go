package scanner

import (
	"fmt"
	"net"
	"sort"

	"github.com/nireo/port-scanner/utils"
)

var address string

func worker(ports, results chan int, address string) {
	for p := range ports {
		addressWithPort := fmt.Sprintf("%s:%d", address, p)
		conn, err := net.Dial("tcp", addressWithPort)
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
	}
}

// RunScanner initializes the scanner and returns the open ports
func RunScanner(addressString string, portsString string) []int {
	address = addressString

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	if portsString == "" {
		for i := 0; i < cap(ports); i++ {
			go worker(ports, results, address)
		}

		go func() {
			for i := 1; i <= 1024; i++ {
				ports <- i
			}
		}()

		for i := 0; i <= 1024; i++ {
			port := <-results
			if port != 0 {
				openPorts = append(openPorts, port)
			}
		}
	} else {
		formattedPorts := utils.SeparatePortsFromString(portsString)

		for i := 0; i < cap(ports); i++ {
			go worker(ports, results, address)
		}

		go func() {
			for _, value := range formattedPorts {
				ports <- value
			}
		}()

		for i := 0; i < len(formattedPorts); i++ {
			port := <-results
			if port != 0 {
				openPorts = append(openPorts, port)
			}
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)

	return openPorts
}
