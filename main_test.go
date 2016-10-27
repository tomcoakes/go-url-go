package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"time"
	"strings"
)

func TestFetchsManyUrls(t *testing.T) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "A")
	}))
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "B")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "C")
	}))

	actualResponse, _ := fetchUrls([]string{ts1.URL, ts2.URL, ts3.URL})

	megaResponse := ""
	for _, resp := range actualResponse {
		megaResponse += resp.respBody
	}

	for _, expected := range([]string{"A","B","C"}) {
		if !strings.Contains(megaResponse,expected) {
			t.Error("holy shit!")
		}
	}
}

func TestConcatenatingUrls(t *testing.T) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "A")
	}))
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "B")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "C")
	}))

	megaResponse, _ := Concatenator(ts1.URL, ts2.URL, ts3.URL)

	for _, expected := range([]string{"A","B","C"}) {
		if !strings.Contains(megaResponse,expected) {
			t.Error("holy shit!")
		}
	}
}

func BenchmarkConcatenator(b *testing.B) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 200)
		fmt.Fprint(w, "Response for first")
	}))
	count := 100
	urls := make([]string, count)
	for i := 0; i < count; i++ {
		urls[i] = ts1.URL
	}
	b.ResetTimer()


	for i := 0; i < b.N; i++ {
		Concatenator(urls...)
	}
}
