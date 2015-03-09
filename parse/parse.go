package parse

import "fmt"

var b_end byte = 0xf5
var b_start byte = 0xfa

func main() {
	fmt.Println("parse")
}

func Ignorenulls(b []byte) []byte {
	var c []byte
	for i := range b {
		c = append(c, b[i])
		if b[i] == b_end {
			break
		}
	}
	return c
}

func Body(b []byte) []byte {
	var c []byte
	c = b[6 : len(b)-2]
	return c

}

func CRC(b2 []byte) byte {
	fmt.Printf("CRCGA keldi %x ", b2)
	var crc byte
	for i := range b2 {
		crc = crc ^ b2[i]
	}
	fmt.Printf("crc %X ", crc)
	return crc
}
