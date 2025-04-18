package main

import (
	"os"
	"time"

	"github.com/julianzillner/DynDNS/request"
)



func main() {
	var url = os.Getenv("URL")
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