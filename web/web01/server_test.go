package web01

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayerScore(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/player/wyett", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)
	t.Run("return wyett's score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("want '%s', get '%s'", want, got)
		}
	})
}
