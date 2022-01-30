package main

import (
	"fmt"
	"net/http"
)

func main() {
	requestUrl := "http://www.baidu.com"
	reponse, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("Err:", err)
	}
	defer reponse.Body.Close()
	fmt.Println(reponse.Body)
}
