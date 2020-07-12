package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe here!"), boring("Ann here!"))
	for i := 0; i < 10; i++ {
		// fmt.Printf("You say: %q\n", <-c)
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func fanIn(input_one, input_two <-chan string) <-chan string {
	c := make(chan string)
	go func() { for {c <- <-input_one} }()
	go func() { for {c <- <-input_two} }()
	return c
}

func boring(msg string) <- chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}() // don't forget about the function call!
	return c
}
