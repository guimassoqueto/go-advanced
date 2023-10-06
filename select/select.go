package main

import (
	"fmt"
	"math/rand"
	"github.com/google/uuid"
)

func main() {
	channelInt := make(chan int)
	channelString := make(chan string)

	go func() {
		channelInt <- 1 + rand.Intn(10)
	}()

	go func() {
		channelString <- uuid.New().String()
	}()

	// only one will be selected if receives message from two channels
	select {
	case intMessage := <- channelInt:
		fmt.Printf("Received a INTEGER message: %d\n", intMessage)
	case stringMessage := <- channelString:
		fmt.Printf("Received a STRING message: %s\n", stringMessage)
	}
}