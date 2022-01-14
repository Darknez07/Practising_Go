package main
import (
	"fmt"
	"os"
	"net"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		for err != nil {
			continue
		}
		//Go routine
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	// Close connection on exit
	defer conn.Close()
	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		if err != nil{
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil{
			return
		}
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error aaya %s",err.Error())
		os.Exit(1)
	}
}