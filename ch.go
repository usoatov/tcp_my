package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		for i := range c {
			time.Sleep(time.Second)
			fmt.Println("i=", i, c)
		}
	}()

	go func() {
		c <- 3
		c <- 1
		c <- 4
		c <- 2

	}()
	time.Sleep(time.Second)
	close(c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	time.Sleep(5 * time.Second)

}
