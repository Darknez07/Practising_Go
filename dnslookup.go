package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s hostname\n", os.Args[0])
		fmt.Println("Usage", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip6", name)
	addr2, err2 := net.ResolveIPAddr("ip4", name)
	if err != nil && err2 != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address IPV6 is: ", addr.String())
	fmt.Println("Resolved address IPV4 is: ", addr2.String())
	network := net.ParseIP(addr2.String())
	mask := network.DefaultMask()
	network = network.Mask(mask)
	fmt.Println("Network address is: ", network.String())
	/* names := os.Args[1]
	addrs, err := net.LookupHost(names)
	if err != nil{
		fmt.Println("Lookup Error", err.Error())
		os.Exit(1)
	}

	for _, addr := range addrs{
		fmt.Println(addr)
	} */

	os.Exit(0)
}
