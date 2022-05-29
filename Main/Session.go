package main

import (
	"math/rand"
	"time"
)

var sessionMap map[string]any

func Init() {
	sessionMap = map[string]any{}
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RegisterSession(s string, data any) {
	sessionMap[s] = data
}

func ReadSession(s string) any {
	return sessionMap[s]
}

func DeleteSession(s string) {
	sessionMap[s] = nil
}

func ModifySession(s string, data any) {
	sessionMap[s] = data
}

func CheckSessionExist(s string) bool {
	return sessionMap[s] != nil
}
