package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julianzillner/DynDNS/request"
	endpoints "github.com/julianzillner/DynDNS/utils"
)



func main() {
	var url = os.Getenv("URL")
	intervalStr := os.Getenv("INTERVAL")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		panic("Invalid INTERVAL value: " + err.Error())
	}

	    go func() {
        http.HandleFunc("/health", endpoints.GetHealth)
        err := http.ListenAndServe(":3333", nil)
        if err != nil {
            panic("Failed to start HTTP server: " + err.Error())
        }
    }()


	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop() 

	for {
		select {
		case <-ticker.C:
			request.Call(url)
		}
	}
}