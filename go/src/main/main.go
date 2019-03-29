package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	message := make(chan string)
	stop := make(chan bool)
	go sender(message)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Time's up!")
		stop <- true
	}()
	for {
		select {
		case <-stop: //通道关闭
			return
		case msg := <-message:
			fmt.Println(msg)
		}
	}
}
