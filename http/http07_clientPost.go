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
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName"
	client := http.Client{}

	param := &url.Values{"theCityName": {"北京"}}
	requestBody := bytes.NewBufferString(param.Encode())

	post, err := client.Post(urlStr, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer post.Body.Close()

	if post.StatusCode == 200 {
		body, err := ioutil.ReadAll(post.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("%+v\n", post)

}
