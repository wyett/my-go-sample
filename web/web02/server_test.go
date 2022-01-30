package web02

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayerStatus(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/player/wyett", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	t.Run("/player/", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if want != got {
			fmt.Printf("want '%s', got '%s'", want, got)
		}
	})
}
