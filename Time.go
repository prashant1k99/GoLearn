package main

import (
	"fmt"
	"time"
)

// Go offers extensive support for times and durations; here are some examples.

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)
	// 2024-08-19 20:56:25.679406 +0530 IST m=+0.000098793

	// You can buld a time struct by providing the year, month, day, etc. Times are always associated with a  Location, i.e. timezone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)
	p(then)
	// 2009-11-17 20:34:58.651387237 +0000 UTC

	// You can extract the various components of the time value as expected.
	p(then.Year())
	// 2009
	p(then.Month())
	// November
	p(then.Day())
	// 17
	p(then.Hour())
	// 20
	p(then.Minute())
	// 34
	p(then.Second())
	// 58
	p(then.Nanosecond())
	// 651387237
	p(then.Location())
	// UTC

	// The Monday-Sunday Weekdays are also available.
	p(then.Weekday())
	// Tuesday

	// These methods compare two times, testing if the first occur before, after, or at the same time as the second, respectively.
	p(then.Before(now))
	// true
	p(then.After(now))
	// false
	p(then.Equal(now))
	// false

	// The Sub methods returns a Duration representing teh interval between two times.
	diff := now.Sub(then)
	p(diff)
	// 129330h54m55.854280763s

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	// 129330.92314752132
	p(diff.Minutes())
	// 7.759855388851279e+06
	p(diff.Seconds())
	// 4.6559132333107674e+08
	p(diff.Nanoseconds())
	// 465591323331076763

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	p(then.Add(diff))
	// 2024-08-19 15:32:47.704227 +0000 UTC
	p(then.Add(-diff))
	// 1995-02-16 01:37:09.598547474 +0000 UTC
}
