package endpoints

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type ipResponse struct {
	Ip string `json:"ip"`
}

func Initialize() {
	
	staticDir := filepath.Join(getWorkingDir(), "static")
    fs := http.FileServer(http.Dir(staticDir))
    http.Handle("/", http.StripPrefix("/", fs))


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

func getWorkingDir() string {
    dir, err := os.Getwd()
    if err != nil {
        panic("Failed to get working directory: " + err.Error())
    }
    return dir
}