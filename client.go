package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func post(url string, data string) {}

func main() {
	method, url := os.Args[1], os.Args[2]
	if method == "GET" {
		get(url) // calls get func
	} else if method == "POST" {
		data := os.Args[3]
		post(url, data) // calls post func
	}
}
