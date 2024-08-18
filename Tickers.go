package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.
// here's an example of a ticker taht ticks periodically utill we stop it.

func main() {
	// Tickers use a similar mechanism to timers: a channel taht is sent values.
	// Here we'll use the select builtin on the channel to await the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// Tick at 2024-08-18 14:52:23.562127 +0530 IST m=+0.501268168
	// Tick at 2024-08-18 14:52:24.062116 +0530 IST m=+1.001253376
	// Tick at 2024-08-18 14:52:24.562107 +0530 IST m=+1.501241584
	// Ticker stopped

	// Tickers can be stopped like timers. Once a ticker is stopped it won;t receive any more values on its channel.
	// We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}