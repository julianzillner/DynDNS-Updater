package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	url := "https://julianzillner.com"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()	


	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	//data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request successful")
	//fmt.Println(string(data))
}