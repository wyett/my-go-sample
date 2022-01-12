package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	strUrl := "http://localhost:8080/demo"
	param := url.Values{"name": {"wyett"}}
	requestBody := bytes.NewBufferString(param.Encode())

	post, err := http.Post(strUrl, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer post.Body.Close()

	if post.StatusCode == 200 {
		postBody, err := ioutil.ReadAll(post.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(postBody))
	}
	fmt.Printf("%+v\n", post)

}
