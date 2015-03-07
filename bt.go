package main

import "fmt"

func main() {
	fmt.Println("bt")
	bt := 0x00ad
	fmt.Printf("%x", bt)
	fmt.Println("")
	hi := bt >> 8
	lo := bt & 0xFF
	fmt.Printf("%x, %x", hi, lo)
	fmt.Println("")

	for i := 0xc0; i <= 0x7fff; i = i + 0x08 {
		hi := i >> 8
		lo := i & 0xFF
		fmt.Printf("H %x L %x, ", hi, lo)

	}
}
