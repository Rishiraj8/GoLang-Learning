package main

import (
    "fmt"
    "time"
)

func ping(ch chan string) {
    for {
        ch <- "Ping" // Send "Ping" into the channel
        time.Sleep(500 * time.Millisecond)
    }
}

func pong(ch chan string) {
    for {
        ch <- "Pong" // Send "Pong" into the channel
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    ch := make(chan string) // Create a channel

    // Start two goroutines
    go ping(ch)
    go pong(ch)

    // Continuously receive and print messages from the channel
    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
}
