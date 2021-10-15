// @author     : wyettLei
// @date       : Created in 2021/4/27 15:24
// @description: TODO

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AlertMsg struct {
}

//func AlertHandler(h http.Handler) Handler {
//    return &alertHandler{}
//}

//type alertHandler struct {
//alertmsg chan(AlertMsg 100)
//AlertMsg
//}

//func (hs *alertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//    //fmt.Println(hs, "hellostruct")
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
//}

func HandleAlertMsg(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(hs, "hellostruct")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, " >", "  "); err != nil {
		panic(err)
	}
	log.Println(buf.String())
}
