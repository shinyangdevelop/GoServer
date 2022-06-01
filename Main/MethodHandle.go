package main

import (
	"io/ioutil"
	"log"
	"net/http"
	template2 "text/template"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template := template2.New("root")
	content, err := ioutil.ReadFile("res/index.html")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	template, err = template.Parse(string(content))
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	var tempData = RootData{
		IsLoggedIn: false,
		UserName:   "",
	}
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		log.Println("No session cookie")
	} else {
		if CheckSessionExist(cookie.Value) {
			tempData = ReadSession(cookie.Value).(RootData)
		}
	}
	err = template.Execute(w, tempData)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func postRoot(w http.ResponseWriter, r *http.Request) {

}

func putRoot(w http.ResponseWriter, r *http.Request) {
}

func deleteRoot(w http.ResponseWriter, r *http.Request) {
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	temp := RandStringRunes(256)
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   temp,
		Expires: time.Now().Add(time.Second * 10),
	})
	RegisterSession(temp, RootData{
		IsLoggedIn: true,
		UserName:   "Steve",
	})
}

func postLogin(w http.ResponseWriter, r *http.Request) {

}

func putLogin(w http.ResponseWriter, r *http.Request) {

}

func deleteLogin(w http.ResponseWriter, r *http.Request) {

}

func getLogout(w http.ResponseWriter, r *http.Request) {
}

func postLogout(w http.ResponseWriter, r *http.Request) {

}

func putLogout(w http.ResponseWriter, r *http.Request) {

}

func deleteLogout(w http.ResponseWriter, r *http.Request) {

}
