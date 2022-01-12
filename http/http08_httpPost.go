package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	post, err := http.Post("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		"application/x-www-form-urlencoded", strings.NewReader("theCityName=北京"))
	if err != nil {
		log.Fatal(err)
	}
	defer post.Body.Close()

	all, err := ioutil.ReadAll(post.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(all))
}
