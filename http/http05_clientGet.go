package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlStr := "http://www.baidu.com"
	client := http.Client{}

	response, err := client.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
}
