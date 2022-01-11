package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"
	response, err := http.Get(urlStr)
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

	fmt.Println("--------------------------------------------")
	fmt.Printf("response: %+v\n", response)
	fmt.Println("--------------------------------------------")
	fmt.Printf("Response.Header %+v\n", response.Header)
	fmt.Printf("Response.codeStatus: %+v\n", response.StatusCode)
	fmt.Printf("Response.Status: %+v\n", response.Status)
	fmt.Printf("Response.body: %+v\n", response.Body)
	fmt.Printf("Response.Cookie: %+v\n", response.Cookies())
	fmt.Printf("Response.Request: %+v\n", response.Request)

}
