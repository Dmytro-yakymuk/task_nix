package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url = "https://jsonplaceholder.typicode.com"
)

func main() {
	resp, err := http.Get(url + "/posts")
	if err != nil {
		log.Fatalf("Error GET request in address %s: %s", url+"/posts", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	fmt.Printf("%s\n", body)
}
