package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resp, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", resp)

}
