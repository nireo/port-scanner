package main

import (
	"flag"
	"fmt"
	"net"
	"sort"

	"github.com/nireo/port-scanner/utils"
)

var address string

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", address, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
	}
}

func main() {
	var portsString string
	flag.StringVar(&portsString, "ports", "", "enter specific ports to be scanned separated by commans i.e: 1,8008,3001")

	var addressFlag string
	flag.StringVar(&addressFlag, "address", "", "enter the address you want to scan")

	if addressFlag == "" {
		fmt.Println("ERROR: You need to provide an address!")
		return
	}
	address = addressFlag

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	if portsString == "" {
		for i := 0; i < cap(ports); i++ {
			go worker(ports, results)
		}

		go func() {
			for i := 1; i <= 1024; i++ {
				ports <- i
			}
		}()

		for i := 0; i < 1024; i++ {
			port := <-results
			if port != 0 {
				openPorts = append(openPorts, port)
			}
		}

	} else {
		formattedPorts := utils.SeparatePortsFromString(portsString)

		for i := 0; i < cap(ports); i++ {
			go worker(ports, results)
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
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
