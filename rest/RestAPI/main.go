package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "select an operation")
}

func add(w http.ResponseWriter, req *http.Request) {
	x := req.URL.Query()["x"][0]
	y := req.URL.Query()["y"][0]
	resp, _ := http.Get("http://localhost:9999?x=" + x + "&y=" + y)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func sub(w http.ResponseWriter, req *http.Request) {
	x := req.URL.Query()["x"][0]
	y := req.URL.Query()["y"][0]
	resp, _ := http.Get("http://localhost:9998?x=" + x + "&y=" + y)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func div(w http.ResponseWriter, req *http.Request) {
	x := req.URL.Query()["x"][0]
	y := req.URL.Query()["y"][0]
	resp, _ := http.Get("http://localhost:9997?x=" + x + "&y=" + y)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func mul(w http.ResponseWriter, req *http.Request) {
	x := req.URL.Query()["x"][0]
	y := req.URL.Query()["y"][0]
	resp, _ := http.Get("http://localhost:9996?x=" + x + "&y=" + y)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/div", div)
	http.HandleFunc("/mul", mul)

	err := http.ListenAndServe("localhost:8888", nil)

	if err != nil {
		log.Fatal(err)
	}
}
