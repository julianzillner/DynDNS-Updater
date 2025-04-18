package main

import (
	"time"

	"github.com/julianzillner/DynDNS/request"
)


func main() {
	var url = "https://julianzillner.com"
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop() 

	for {
		select {
		case <-ticker.C:
			request.Call(url)
		}
	}
}