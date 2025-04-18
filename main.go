package main

import (
	"time"

	"github.com/julianzillner/DynDNS/request"
)



func main() {
	var url = "https://julianzillner.com"
	var interval int = 1


	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop() 

	for {
		select {
		case <-ticker.C:
			request.Call(url)
		}
	}
}