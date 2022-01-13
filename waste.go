package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	service := ":8010"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleConn(conn)
		conn.Close()
	}
}

func handleConn(conn net.Conn){
	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		if err != nil{
			return
		}
		fmt.Printf(" %s ",strings.TrimSpace(string(buf[0:])))
		_, err = conn.Write([]byte(strings.TrimSpace(string(buf[0:n]))))
		if err != nil{
			return
		}
	}
}

func checkError(err error){
	if err !=nil{
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}