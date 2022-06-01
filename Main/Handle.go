package main

import (
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getRoot(w, r)
	case "POST":
		postRoot(w, r)
	case "PUT":
		putRoot(w, r)
	case "DELETE":
		deleteRoot(w, r)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLogin(w, r)
	case "POST":
		postLogin(w, r)
	case "PUT":
		putLogin(w, r)
	case "DELETE":
		deleteLogin(w, r)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLogout(w, r)
	case "POST":
		postLogout(w, r)
	case "PUT":
		putLogout(w, r)
	case "DELETE":
		deleteLogout(w, r)
	}
}
