// @author     : wyettLei
// @date       : Created in 2021/4/26 11:31
// @description: TODO

package httpsample

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func HttpTimeoutMain() {
	srv := http.Server{
		Addr: ":8888",
		//Handler:           http.HandlerFunc(TimeoutHandleFunc),
		Handler:      http.TimeoutHandler(http.HandlerFunc(TimeoutHandleFunc), 1*time.Second, "timeout!\n"),
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func TimeoutHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("TimeoutHandleFunc()...")
	time.Sleep(2 * time.Second)
	io.WriteString(w, "time out handler")
}

//--------------------------------------------------------------------------

func slowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("slowHandler()...")
	result := callSlowApi()
	io.WriteString(w, result+"\n")
}

func callSlowApi() string {
	t := rand.Intn(5)
	select {
	case <-time.After(time.Duration(t) * time.Second):
		log.Printf("wait for %d seconds\n", t)
		return "slow over"
	}
}

func HttpSlowMain() {
	srv := http.Server{
		Addr: ":8888",
		//Handler:           http.HandlerFunc(TimeoutHandleFunc),
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "timeout!\n"),
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

//----------------------------------------------------------------------

func callContextSlowApi(ctx context.Context) string {
	t := rand.Intn(5)
	select {
	case <-ctx.Done():
		{
			log.Printf("call slow api was interputted, time was %d seconds", t)
			return "interruptted"
		}
	case <-time.After(time.Duration(t) * time.Second):
		{
			log.Printf("call slow api after %d seconds", t)
			return "slow over"
		}

	}
}

func slowContextHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("slowContextHandler()...")
	result := callContextSlowApi(r.Context())
	io.WriteString(w, result+"\n")
}

func HttpSlowContextMain() {
	srv := http.Server{
		Addr:         ":8888",
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowContextHandler), 1*time.Second, "timeout!\n"),
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
