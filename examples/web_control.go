package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"vba_control"
)

var vba vba_control.Client

const (
	keypath = "/button/"
)

func keyHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(keypath):]

	vba.SendInput(strings.TrimSpace(key))
	fmt.Println("pressed", key)
	fmt.Fprint(w, key)
}

func main() {
	rom := os.Args[1]
	vba = *vba_control.New(rom)

	go vba.Start()
	fmt.Println("Starting VBA with", rom)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc(keypath, keyHandler)

	http.ListenAndServe(":8080", nil)
}
