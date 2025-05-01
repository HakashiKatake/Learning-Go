package main

import (
	"fmt"
	"time"
)

// send
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number:", num)
		time.Sleep(time.Second * 1) // Simulate some processing time
	}

}

// receive
func sum(result chan int, a int, b int) {
	numresult := a + b
	result <- numresult
}

// blocking in channels
func send(ch chan bool) {
	defer func() { ch <- true }() //block until a receiver is available
	fmt.Println("Sending message...")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Close the channel when done
	for email := range emailChan {
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second) // Simulate some processing time
	}
}

func main() {
	// Channels
	// Channels are used to communicate between goroutines.
	// They are a way to send and receive data between goroutines.
	// A channel is created using the `make` function.
	// Example:
	// ch := make(chan int)
	// ch <- 1 // Send 1 to the channel
	// val := <-ch // Receive from the channel
	// fmt.Println(val) // Print the value received from the channel

	// messageChan := make(chan string)

	// messageChan <- "Ping" //block until a receiver is available

	// msg := <-messageChan //block until a sender is available

	// fmt.Println(msg)

	// result := make(chan int)

	// go sum(result, 4, 5)
	// numresult := <-result
	// fmt.Println("Sum:", numresult)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//blocking in channels
	// ch := make(chan bool)
	// go send(ch)
	// <-ch //block until a sender is available

	emailChan := make(chan string, 100) // Buffered channel
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done sending.")

	close(emailChan) // Close the channel to signal completion
	<-done

}
