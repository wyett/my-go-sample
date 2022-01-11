package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"
	// client
	client := http.Client{}

	// request
	request, err := http.NewRequest("Get", urlStr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// doRequest
	response, err := client.Do(request)
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

	fmt.Println("-------------------------------------")
	fmt.Printf("Response：%+v\n", response)
	fmt.Println("-------------------------------------")
	fmt.Printf("Response.Header %+v\n", response.Header)
	fmt.Printf("Response.Status %+v\n", response.Status)
	fmt.Printf("Response.StatusCode %+v\n", response.StatusCode)
	fmt.Printf("Response.Cookies %+v\n", response.Cookies())
	fmt.Printf("Response.Request %+v\n", response.Request)
	fmt.Printf("Response.Body %+v\n", response.Body)

}
