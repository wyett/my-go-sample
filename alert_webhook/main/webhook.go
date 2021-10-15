// @author     : wyettLei
// @date       : Created in 2021/4/27 9:33
// @description: TODO

package main

import (
	"log"
	"net/http"
	"time"
)

//func AlertMsgHandler(w http.ResponseWriter, r *http.Request) {
//    b, err := ioutil.ReadAll(r.Body)
//    if err != nil {
//        panic(err)
//    }
//    defer r.Body.Close()
//    var buf bytes.Buffer
//    if err := json.Indent(&buf, b, " >", "  "); err != nil {
//        panic(err)
//    }
//    log.Println(buf.String())
//    fmt.Println(buf.String())
//}

func main() {
	//alertHandler := alertHandler{}

	srv := http.Server{
		Addr:         ":5001",
		Handler:      http.TimeoutHandler(http.HandlerFunc(HandleAlertMsg), 1*time.Second, "time out\n"),
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
