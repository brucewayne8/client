package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func readResponse(resBody io.Reader) { // reads the response body and prints to terminal
	body, err := ioutil.ReadAll(resBody)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func get(url string) { // get func
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	readResponse(res.Body)
}

func post(url string, data string) { //post func
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
	readResponse(res.Body)
}

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		panic("Give the arguments in the  mentioned format")
	}
	method, url := os.Args[1], os.Args[2]
	if method == "GET" {
		get(url) // calls get func
	} else if method == "POST" {
		if len(os.Args) != 4 {
			panic("Give the 4 arguments including executable file  for POST call")
		}
		data := os.Args[3]
		post(url, data) // calls post func
	}
}
