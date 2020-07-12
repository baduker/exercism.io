package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	// fmt.Println("I'm listening.")
	// time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <- c) // Receiving from channel; it's just a value
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// fmt.Println(msg, i)
		c <- fmt.Sprintf("%s %d", msg, i)
		// Expression to be sent can be any sutiable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
