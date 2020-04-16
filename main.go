package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/panditb/golearning/config"
)

func main() {
	fmt.Println(" Starting the rest consume application...")
	config := config.GetConfig()
	url := os.Getenv("url")
	if len(url) == 0 {
		fmt.Println("Env URL is null")
		url = config.ConsumedUrl
	} else {
		fmt.Println("Env URL is not null")
	}

	response, error := http.Get(url)
	if error != nil {
		fmt.Printf("The HTTP request failed with error %s\n", error)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	fmt.Println("Terminating the application...")

}
