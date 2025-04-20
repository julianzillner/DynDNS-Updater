package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dimiro1/banner"
	noip "github.com/julianzillner/DynDNS/provider"
	"github.com/julianzillner/DynDNS/request"
	endpoints "github.com/julianzillner/DynDNS/utils"
	"github.com/mattn/go-colorable"
)

func main() {
	var url = os.Getenv("URL")
	intervalStr := os.Getenv("INTERVAL")
	provider := os.Getenv("PROVIDER")

	var noipUsername = os.Getenv("NOIP_USERNAME")
	var noipPassword = os.Getenv("NOIP_PASSWORD")
	var noipHost = os.Getenv("NOIP_HOST")

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		panic("Invalid INTERVAL value: " + err.Error())
	}


	go func() {
		fileServer := http.FileServer(http.Dir("./static"))
     	http.Handle("/", fileServer)
		endpoints.Initialize()
		}()
		

	if provider == "noip" {
		url = noip.Call(noipUsername, noipPassword, noipHost)
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


func init() {
	templ := `{{ .Title "DynDNS Updater" "" 2 }}
   Developer: Julian Zillner 
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
