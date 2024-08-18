package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    client := http.Client{
        Timeout: 2 * time.Second, // Set a 2-second timeout
    }

    resp, err := client.Get("https://example.com") // Make a GET request
    if err != nil {
        fmt.Println("Request failed:", err)
    } else {
        fmt.Println("Request succeeded with status:", resp.Status)
        resp.Body.Close()
    }
}

/* 
Explanation:
HTTP Client with Timeout:
The http.Client is configured with a 2-second timeout.
If the request to https://example.com takes longer than 2 seconds, the client will cancel it and return an error.
Handling the Response:
If the request succeeds within the timeout, the status is printed.
If it takes too long, an error is printed indicating the request failed due to the timeout.
*/