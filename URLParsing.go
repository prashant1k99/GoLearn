package main

import (
	"fmt"
	"net"
	"net/url"
)

// URLs provide a uniform way to locate resources. Here's how to parse URLs in Go.

func main() {
	// We'll parse this exmaple URL, which includes a scheme, authentication info, host, port, path, query params and query fragment.'
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse teh URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)
	// postgres

	// User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println(u.User)
	// user:pass
	fmt.Println(u.User.Username())
	// user
	p, _ := u.User.Password()
	fmt.Println(p)
	// pass

	// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println(u.Host)
	// host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	// host.com
	fmt.Println(port)
	// 5432

	// Here we extract the path and the fragment after the #.
	fmt.Println(u.Path)
	// /path
	fmt.Println(u.Fragment)
	// f

	// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
	fmt.Println(u.RawQuery)
	// k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// map[k:[v]]
	fmt.Println(m["k"][0])
	// v
}
