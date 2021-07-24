package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("tracing start Request URL & method: %s, %s", r.URL, r.Method)
		next.ServeHTTP(rw, r)
		log.Println("tracing end")
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("logging start Request address: %s", r.RemoteAddr)
		next.ServeHTTP(rw, r)
		log.Println("logging end")
	})
}

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("timeRecording start")
		startTime := time.Now()
		next.ServeHTTP(rw, r)
		elapsedTime := time.Since(startTime)
		log.Printf("execution time: %s", elapsedTime)
		log.Println("timeRecording end")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	// TODO: transform middleware functions using Alice
	// From Middleware1(Middleware2(Middleware3(App)))
	// To alice.New(Middleware1, Middleware2, Middleware3).Then(App)
	http.Handle("/", tracing(logging(timeRecording(http.HandlerFunc(hello)))))
	http.ListenAndServe(":8080", nil)
}
