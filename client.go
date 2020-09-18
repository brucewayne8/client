package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func post(url string, data string) {
	parts := strings.Split(data, "&&")
	leng := len(parts)

	if leng == 0 {
		panic("Provide the field values")
	}

	var mp = make(map[string]string)

	for i := 0; i < leng; i++ {
		fieldValue := strings.Split(parts[i], "=")
		if len(fieldValue) != 2 {
			panic("Provide the field values in correct format")
		}
		mp[fieldValue[0]] = fieldValue[1]
	}

	jsonbody, err := json.Marshal(mp)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbody))
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

func main() {
	method, url := os.Args[1], os.Args[2]
	if method == "GET" {
		get(url) // calls get func
	} else if method == "POST" {
		data := os.Args[3]
		post(url, data) // calls post func
	}
}
