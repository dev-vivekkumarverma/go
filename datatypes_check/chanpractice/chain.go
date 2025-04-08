package chanpractice

import (
	"fmt"
	"sync"
	"time"
)

// takes message, send the message to the given channel.
func SendMessage(message string, channel chan string, wg *sync.WaitGroup, slpt int) {
	defer wg.Done()
	fmt.Println("sending message to the channel...")
	time.Sleep(time.Duration(slpt) * time.Second)
	channel <- message

	fmt.Println("message sent to the channel.")
}

func ReceiveMessage(channel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-channel
	if msg != "" {
		fmt.Println("Message recieved !")
		fmt.Println("Message::", msg)
	} else {
		fmt.Println("Message not received yet")
	}
}
