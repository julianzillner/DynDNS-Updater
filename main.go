package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julianzillner/DynDNS/request"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}


func main() {
	godotenv.Load()
	var url = os.Getenv("URL")
	intervalStr := os.Getenv("INTERVAL")
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		panic("Invalid INTERVAL value: " + err.Error())
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop() 

	for {
		select {
		case <-ticker.C:
			request.Call(url)
		}
	}
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)
}