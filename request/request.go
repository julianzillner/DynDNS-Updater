package request

import (
	"log"
	"net/http"

	"github.com/fatih/color"
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

	//data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	color.Green("Request successful")
	//fmt.Println(string(data))
}