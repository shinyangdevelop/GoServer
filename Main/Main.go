package main

import "net/http"

func main() {
	Init()
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":80", nil)
}
