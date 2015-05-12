package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/usoatov/tcp_my/parse"
)

func Send(s []byte) []byte {
	servAddr := "84.204.80.186:1990"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
	}

	_, err = conn.Write(s)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", s)

	reply := make([]byte, 150)

	_, err = conn.Read(reply)

	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	conn.Close()
	time.Sleep(1100 * time.Millisecond)

	return reply

}

func main() {
	var by []byte

	for i := 0xc0; i <= 0x7fff; i = i + 0x78 {
		fmt.Printf("\nCikl i=%x ", i)

		hi := byte(i >> 8)
		lo := byte(i & 0xFF)
		fmt.Printf("H %x L %x, ", hi, lo)
		by = []byte{0xfa, 0x01, 0x01, 0x04, 0x00, 0xa0}
		by = append(by, 0x78, hi, lo)

		crc := parse.CRC(by[3:len(by)])
		by = append(by, crc, 0xf5)

		for j := range by {
			fmt.Printf("%X; ", by[j])
		}
		fmt.Println("------------------")

		res := Send(by)

		//, 0x78, 0x00, 0xc0, 0x1c, 0xf5}

		r := parse.Ignorenulls(res)

		println("reply from server=", string(r))

		for i := range r {
			fmt.Printf("%x ", r[i])
		}

		fmt.Println("")
		bd := parse.Body(r)
		fmt.Printf("Body %x ", bd)
		fmt.Println("------------------")

		for j := 0; j <= 14; j = j + 1 {
			fmt.Print("key ")
			fmt.Printf("%X ", bd[j*8])
			fmt.Printf("%X ", bd[j*8+1])
			fmt.Printf("%X ", bd[j*8+2])
			fmt.Printf("%X ", bd[j*8+3])
			fmt.Printf("%X ", bd[j*8+4])
			fmt.Printf("%X ", bd[j*8+5])
			fmt.Printf("%X ", bd[j*8+6])
			fmt.Printf("%X ", bd[j*8+7])
			fmt.Println("")
		}

	}

}
