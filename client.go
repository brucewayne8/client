package main

import (
	"os"
)

func get(url string) {}

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
