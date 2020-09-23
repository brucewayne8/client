package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(url string) { // get func
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("check thr url ", err)
	}
	defer res.Body.Close()
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal("Error while copying response to stdout ", err)
	}
}

func post(url string, data string, flag int) { //post func // flag for data or file

	var jsonbody []byte
	var err error

	if flag == 0 {
		parts := strings.Split(data, "&&")
		leng := len(parts)

		if leng == 0 {
			log.Fatal("Provide the field values")
		}

		var mp = make(map[string]string)

		for i := 0; i < leng; i++ {
			fieldValue := strings.Split(parts[i], "=")
			if len(fieldValue) != 2 {
				log.Fatal("Provide the field values in correct format")
			}
			mp[fieldValue[0]] = fieldValue[1]
		}

		jsonbody, err = json.Marshal(mp)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		jsonbody, err = ioutil.ReadFile(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonbody))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal("Error while copying response to stdout ", err)
	}
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 5 {
		log.Fatal("Give the arguments in the  mentioned format")
	}
	method, url := os.Args[1], os.Args[2]
	if method == "GET" {
		get(url) // calls get func
	} else if method == "POST" {
		if len(os.Args) < 4 {
			log.Fatal("Provide arguments in the mentioned format")
		}
		if len(os.Args) == 4 {
			data := os.Args[3]
			post(url, data, 0) // calls post func
		} else if os.Args[3] == "FILE" {
			data := os.Args[4]
			post(url, data, 1)
		} else {
			log.Fatal("Give the arguments in the  mentioned format")
		}
	}
}
