package parse

import "fmt"

var b_end byte = 0xf5

func main() {
	fmt.Println("parse")
}

func Ignorenulls(b []byte) []byte {
	var c []byte
	for i := range b {
		//fmt.Println(b[i])
		c = append(c, b[i])
		if b[i] == b_end {
			break
		}
	}
	return c
}
