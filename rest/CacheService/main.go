package main

import (
	"fmt"
	"log"
	"net/http"
)

type cacheService struct {
	results map[key]string
}

type key struct {
	op, x, y string
}

func (localStorage cacheService) postToCache(requestObject key, res string) {
	localStorage.results[requestObject] = res
}

func (localStorage cacheService) getFromCache(requestObject key) string {
	value, ok := localStorage.results[requestObject]
	if !ok {
		return ""
	} else {
		return value
	}
}

func (localStorage cacheService) deleteFromCache(requestObject key) {
	_, ok := localStorage.results[requestObject]
	if ok {
		delete(localStorage.results, requestObject)
	} else {
		return
	}
}

func (localStorage cacheService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query()["op"][0]
	x := r.URL.Query()["x"][0]
	y := r.URL.Query()["y"][0]

	requestObject := key{op: op, x: x, y: y}

	switch r.Method {
	case "POST":
		res := r.URL.Query()["res"][0]
		localStorage.postToCache(requestObject, res)
	case "GET":
		result := localStorage.getFromCache(requestObject)
		fmt.Fprintf(w, result)
	case "PUT":
	case "DELETE":
		localStorage.deleteFromCache(requestObject)
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}

func main() {
	localStorage := &cacheService{results: map[key]string{}}

	err := http.ListenAndServe("localhost:7777", localStorage)

	if err != nil {
		log.Fatal(err)
	}
}
