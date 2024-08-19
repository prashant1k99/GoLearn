package main

import (
	"fmt"
	"time"
)

func main() {
	rate := time.Second // 1 request per second
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	for i := 0; i < 20; i++ {
		<-ticker.C // Wait for the next tick
		fmt.Printf("Request %d processed at %s\n", i+1, time.Now())
	}
}

/*
Explanation:
time.Ticker: This creates a ticker that sends a signal (on its channel C) at regular intervals. In this case, once per second.

<-ticker.C: This blocks until the next tick is received, ensuring that requests are processed at the desired rate (one request per second).

Request 1 processed at 2024-08-19 11:36:52.656655 +0530 IST m=+1.001190210
Request 2 processed at 2024-08-19 11:36:53.656727 +0530 IST m=+2.001252460
Request 3 processed at 2024-08-19 11:36:54.656691 +0530 IST m=+3.001207293
Request 4 processed at 2024-08-19 11:36:55.656645 +0530 IST m=+4.001151210
Request 5 processed at 2024-08-19 11:36:56.656708 +0530 IST m=+5.001204126
Request 6 processed at 2024-08-19 11:36:57.656697 +0530 IST m=+6.001183585
Request 7 processed at 2024-08-19 11:36:58.656726 +0530 IST m=+7.001202960
Request 8 processed at 2024-08-19 11:36:59.656757 +0530 IST m=+8.001224168
Request 9 processed at 2024-08-19 11:37:00.656258 +0530 IST m=+9.000715960
Request 10 processed at 2024-08-19 11:37:01.65756 +0530 IST m=+10.002007626
Request 11 processed at 2024-08-19 11:37:02.656762 +0530 IST m=+11.001200085
Request 12 processed at 2024-08-19 11:37:03.656765 +0530 IST m=+12.001193501
Request 13 processed at 2024-08-19 11:37:04.656767 +0530 IST m=+13.001186043
Request 14 processed at 2024-08-19 11:37:05.65656 +0530 IST m=+14.000968543
Request 15 processed at 2024-08-19 11:37:06.656287 +0530 IST m=+15.000686585
Request 16 processed at 2024-08-19 11:37:07.65679 +0530 IST m=+16.001179460
Request 17 processed at 2024-08-19 11:37:08.656826 +0530 IST m=+17.001206335
Request 18 processed at 2024-08-19 11:37:09.656814 +0530 IST m=+18.001183668
Request 19 processed at 2024-08-19 11:37:10.656821 +0530 IST m=+19.001181751
Request 20 processed at 2024-08-19 11:37:11.656085 +0530 IST m=+20.000436126
*/