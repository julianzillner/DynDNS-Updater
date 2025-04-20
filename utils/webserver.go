package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ipResponse struct {
	Ip string `json:"ip"`
}

func Initialize() {
	http.HandleFunc("/interval/sec", func(w http.ResponseWriter, r *http.Request) {
		GetInterval(w, r, "sec")
	})

	http.HandleFunc("/interval/min", func(w http.ResponseWriter, r *http.Request) {
		GetInterval(w, r, "min")
	})



  	http.HandleFunc("/health", GetHealth)
	http.HandleFunc("/currentIP", GetIpAdress)
    err := http.ListenAndServe(":3333", nil)
    if err != nil {
        panic("Failed to start HTTP server: " + err.Error())
    }
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Healthy\n")
}

func GetIpAdress(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	 resp, _ := http.Get("https://api.ipquery.io/?format=json")
	 body, _ := ioutil.ReadAll(resp.Body)

	 ipResponse := ipResponse{}
	 err := json.Unmarshal(body, &ipResponse)
	 if err != nil {
		 http.Error(w, "Failed to parse IP response", http.StatusInternalServerError)
		 return
	 }
	 io.WriteString(w, ipResponse.Ip)
}

func GetInterval(w http.ResponseWriter, r * http.Request, outputType string) {
	w.WriteHeader(http.StatusOK)
	if outputType == "sec" {
		io.WriteString(w, os.Getenv("INTERVAL"))
	}  else if outputType == "min" {
	intervalStr := os.Getenv("INTERVAL")
	if interval, err := strconv.Atoi(intervalStr); err != nil {
		http.Error(w, "Invalid INTERVAL value", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "%.2f", float64(interval)/100.0)
	}
}
}