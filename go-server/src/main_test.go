package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Setenv("FEELING", "curious")
	t.Setenv("AWESOMNESS", "snacks")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Howdy! I feeeeel curious today. I will tell you my secret of awesomeness: \"snacks\"."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloHandlerDefaultMessage(t *testing.T) {
	os.Unsetenv("FEELING")
	os.Unsetenv("AWESOMNESS")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	expected := "Howdy! I feeeeel $FEELING today. I will tell you my secret of awesomeness: \"$AWESOMNESS\"."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected default body: got %v want %v", rr.Body.String(), expected)
	}
}
