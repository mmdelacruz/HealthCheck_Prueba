package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/pong", nil)
	w := httptest.NewRecorder()

	ping(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	body := w.Body.String()
	expected := "pong"

	if body != expected {
		t.Errorf("expected body %q; got %q", expected, body)
	}
}
