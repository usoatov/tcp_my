package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	servAddr := "84.204.80.186:1990"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	by := []byte{0xfa, 0x01, 0x01, 0x04, 0x00, 0xa0, 0x08, 0x00, 0xaa, 0x06, 0xf5}
	_, err = conn.Write(by)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", by)

	reply := make([]byte, 200)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))
	for i := range reply {
		fmt.Printf("%x ", reply[i])
	}

	conn.Close()
}
