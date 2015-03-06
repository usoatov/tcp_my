package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a := 4
	b := 6
	d := 1
	c := a ^ b ^ d
	fmt.Println(c)
	r := "ъ8<х"
	fmt.Println("len", len(r))
	f := []byte(r)
	fmt.Println(f)
	s := hex.Dump(f)
	fmt.Println(s)

	by := []byte{0xfa, 0x01, 0x01, 0x04, 0x00, 0xa0, 0x08, 0x00, 0x01}
	b2 := []byte{0x04, 0x00, 0xa0, 0x78, 0x00, 0xc0}
	fmt.Println(by)
	var crc byte
	for i := range b2 {
		fmt.Println("b2", i, "=", b2[i])
		crc = crc ^ b2[i]
		fmt.Println("crc", crc)

	}
	//fmt.Println(hex.EncodeToString(crc))

	j := "ъUUх"
	f = []byte(j)
	fmt.Println(f)
	s = hex.Dump(f)
	fmt.Println(s)
}
