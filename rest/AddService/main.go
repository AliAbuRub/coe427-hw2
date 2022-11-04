package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Model the service with an interface
// Capital 'S' in the identifier means the interface is public
type AddService interface {
	PostSum(x int, y int)
	GetSum(x int, y int) (sum int)
	PutSum(x int, y int)
	DeleteSum(x int, y int)
}

func PostSum(x string, y string) {
	//Perform addition and store sum in the cache
	num1, _ := strconv.Atoi(x)
	num2, _ := strconv.Atoi(y)
	sum := num1 + num2

	res := strconv.Itoa(sum)
	_, err := http.Post("http://localhost:7777/?op=add&x="+x+"&y="+y+"&res="+res, "application/json", nil)
	if err != nil {
		log.Fatal()
	}

}

func GetSum(x string, y string) (sum string) {

	resp, err := http.Get("http://localhost:7777/?op=add&x=" + x + "&y=" + y)
	if err != nil {
		log.Fatal()
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if len(body) == 0 { //if length is 0 then there is no entry in the cache
		sum = ""
	} else {
		sum = string(body)
	}
	return
}

func DeleteSum(x string, y string) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", "http://localhost:7777/?op=add&x="+x+"&y="+y, nil)
	_, err := client.Do(req)
	if err != nil {
		log.Fatal()
	}

}

// Make the service available on the network (i.e., callable by other services)
// Function name must be 'ServeHTTP'
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query()["x"][0]
	y := r.URL.Query()["y"][0]

	switch r.Method {
	case "POST":
		PostSum(x, y)
		fmt.Fprintf(w, "New sum saved!")
	case "GET":
		sum := GetSum(x, y)
		if sum == "" {
			fmt.Fprintf(w, "Sum does not exist in cache!")
		} else {
			res, _ := strconv.Atoi(sum)
			fmt.Fprintf(w, "Sum = %v", res)
		}
	case "PUT":
	case "DELETE":
		DeleteSum(x, y)
		fmt.Fprintf(w, "Sum deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}

func main() {

	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe("localhost:9999", nil)
}
