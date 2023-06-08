package main

import "fmt"

type IPAddr []byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	if len(ipAddr) != 4 {
		return "BAD"
	} else {
		return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
	}
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":   {127, 0, 0, 1},
		"googleDNS":  {8, 8, 8, 8},
		"googleDNS2": {8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
