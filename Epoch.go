package main

import (
	"fmt"
	"time"
)

// A common requirement in programs is getting the number of seconds, ms, or ns since the Unix epoch.

func main() {
	// Use time.Now with unix, UnixMilli or UnixNano to get elapsed time since the Unix epoch in seconds, ms or ns, repsectively7
	now := time.Now()
	fmt.Println(now)
	// 2024-08-19 21:06:27.427997 +0530 IST m=+0.000248251

	fmt.Println(now.Unix())
	// 	1724081884
	fmt.Println(now.UnixMilli())
	// 1724081884716
	fmt.Println(now.UnixNano())
	// 1724081884716818000

	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(now.Unix(), 0))
	// 	2024-08-19 21:09:01 +0530 IST
	fmt.Println(time.Unix(0, now.UnixNano()))
	// 2024-08-19 21:09:01.443985 +0530 IST
}
