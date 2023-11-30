package main

import (
	"net/http"
	"sso-go-oauth-demo/backend"
)

func main() {
	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.HandleFunc("/signin", backend.Signin)
	http.HandleFunc("/callback", backend.Callback)
	http.ListenAndServe(":3000", nil)
}
