package main

import (
	"os"
	"strconv"
	"time"

	"github.com/julianzillner/DynDNS/request"
)



func main() {
	var url = os.Getenv("URL")
	intervalStr := os.Getenv("INTERVAL")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		panic("Invalid INTERVAL value: " + err.Error())
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop() 

	for {
		select {
		case <-ticker.C:
			request.Call(url)
		}
	}
}