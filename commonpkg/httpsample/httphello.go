// @author     : wyettLei
// @date       : Created in 2021/4/12 18:00
// @description: TODO

package httpsample

import (
	"fmt"
	"net/http"
)

//--------------------------HandleFunc-----------------------------/
func HttpHello() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8080", nil)
}

//--------------------------HandlerFunc----------------------------------------/

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func HttpHello2() {
	handler := http.HandlerFunc(helloFunc)
	if err := http.ListenAndServe(":8081", handler); err == nil {
		panic(err)
	}
}

//--------------------------------------------------------------/

type hFunc func(w http.ResponseWriter, r *http.Request)

func (h hFunc) hello(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))
	h(w, r)
}

//----------------------------------------------------/
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func HttpHello3() {
	http.HandleFunc("/hello3", hello)
	http.ListenAndServe(":8082", nil)
}

//----------------------------------------------------/

type HelloHandler struct{}

//必须实现ServeHTTP方法
func (hs HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(hs, "hellostruct")
}

func HttpHello4() {
	hs := HelloHandler{}
	http.Handle("/hello4", hs)
	http.ListenAndServe(":8083", nil)
}

//
