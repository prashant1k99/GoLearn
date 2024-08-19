package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of Service.
// Go elegantly supports rate limiting with goroutines, channels and tickers.

func main() {
	// First we'll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We'll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// This limitier channel will receive a value every 200 ms. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receiver from the limiter channel before serving each request, we limit ourselves to 1 request every 200 ms.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// We may want to allow short burts of requests in our rate  limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel. This burstyLimiter channel will aloow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms we'll try to add a new value to burstylimiter, up to its limit of 3.
	go func ()  {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incming requests. The first 3 of these will benefit from teh burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
Output:
request 1 2024-08-19 15:53:27.753618 +0530 IST m=+0.201267376
request 2 2024-08-19 15:53:27.953715 +0530 IST m=+0.401365667
request 3 2024-08-19 15:53:28.153669 +0530 IST m=+0.601321667
request 4 2024-08-19 15:53:28.353638 +0530 IST m=+0.801292876
request 1 2024-08-19 15:53:28.35368 +0530 IST m=+0.801334292
request 2 2024-08-19 15:53:28.353717 +0530 IST m=+0.801372001
request 3 2024-08-19 15:53:28.353722 +0530 IST m=+0.801376667
request 4 2024-08-19 15:53:28.553801 +0530 IST m=+1.001457834
*/