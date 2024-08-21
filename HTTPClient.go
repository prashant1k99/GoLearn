package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
// In this example we'll use it to issue a simple HTTP requests.

func main() {
	// Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method;
	// it uses the http.DefaultClient object which has useful default settings.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP respinse status
	fmt.Println("Response Status:", resp.Status)
	// Response Status: 200 OK

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	// <!DOCTYPE html>
	// <html>
	//   <head>
	//     <meta charset="utf-8">
	//     <title>Go by Example</title>
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
