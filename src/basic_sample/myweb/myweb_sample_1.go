package myweb

import (
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/20 17:42
 * @description: simple sample
 */

func echo(rw http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.Write([]byte("read error"))
		return
	}

	writeLen, writeErr := rw.Write(msg)
	if writeErr != nil || writeLen != len(msg) {
		log.Println("write failed:", writeErr)
	}
}

func WebMain1() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
