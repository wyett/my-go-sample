package web02

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len("/player/"):]

	if s == "wyett" {
		fmt.Fprint(w, 10)
		return
	}

	if s == "boze" {
		fmt.Fprint(w, 20)
	}
}
