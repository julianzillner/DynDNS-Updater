package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dimiro1/banner"
	"github.com/julianzillner/DynDNS/request"
	endpoints "github.com/julianzillner/DynDNS/utils"
	"github.com/mattn/go-colorable"
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
		http.HandleFunc("/currentIP", endpoints.GetIpAdress)
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


func init() {
	templ := `{{ .Title "DynDNS Updater" "" 2 }}
   GoVersion: {{ .GoVersion }}
   GOOS: {{ .GOOS }}
   GOARCH: {{ .GOARCH }}
   NumCPU: {{ .NumCPU }}
   GOROOT: {{ .GOROOT }}
   Compiler: {{ .Compiler }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}

`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
}
