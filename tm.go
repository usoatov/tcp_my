package main

import (
	"fmt"
	"log"
	"time"
)

func doSomething(s string) {
	log.Printf("doing something: %v", s)
}

func startPolling() {
	for _ = range time.Tick(2 * time.Second) {
		doSomething("awesome")
	}

}

func main() {
	msg := make(chan string, 1)
	go startPolling()
	log.Println("this will never printed")
	add :=make(chan int)
	mul :=make(chan int)
	numbs := [1000000]int{}
	go func() {
		total := 0
		for _, v := range numbs {
			total += v
		}
		add <- total
	}()

	go func() {
		total := 1
		for _, v := range numbs {
			total *= v
		}
		mul <- total
	}

	fmt.Println("sum and mul", <-add, <-mul)
}
