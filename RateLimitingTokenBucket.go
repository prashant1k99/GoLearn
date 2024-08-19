package main

import (
	"fmt"
	"time"
)

func main() {
	rate := 1 * time.Second  // Time between tokens
	bucketSize := 3          // Number of tokens that can be accumulated

	tokens := make(chan struct{}, bucketSize)

	// Token refill goroutine
	go func() {
		ticker := time.NewTicker(rate)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case tokens <- struct{}{}:
				default:
					// Bucket is full, discard token
				}
			}
		}
	}()

	for i := 0; i < 5; i++ {
		<-tokens // Take a token
		fmt.Printf("Request %d processed at %s\n", i+1, time.Now())
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

/*
Token Bucket: The tokens channel holds tokens representing available request slots. Tokens are added to the channel at a fixed rate (1 per second) but are limited by the bucket size (3 tokens).

Processing Requests: Each request takes a token from the bucket. If no tokens are available, the request is blocked until a token is added.

Output:
Request 1 processed at 2024-08-19 13:03:51.475476 +0530 IST m=+1.001170501
Request 2 processed at 2024-08-19 13:03:52.474522 +0530 IST m=+2.000215668
Request 3 processed at 2024-08-19 13:03:53.474947 +0530 IST m=+3.000639334
Request 4 processed at 2024-08-19 13:03:54.475511 +0530 IST m=+4.001202376
Request 5 processed at 2024-08-19 13:03:55.47505 +0530 IST m=+5.000740251
*/