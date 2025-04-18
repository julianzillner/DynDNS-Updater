package request

import (
	"log"
	"net/http"
)


func Call(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()	


	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}


	if err != nil {
		log.Fatal(err)
	}
}