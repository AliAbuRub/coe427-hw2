package main

import (
	"fmt"
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

// Define the in-memory (runtime) data store (cache)
// Lower-case letter 's' means the struct is private
type addService struct {
	sums map[key]int
}

type key struct {
	x, y int
}

func (ss *addService) PostSum(x int, y int) {
	//Perform addition and store sum in the cache
	ss.sums[key{x, y}] = x + y
}

func (ss *addService) GetSum(x int, y int) (sum int) {
	//Check if the sum already exists in the cache
	//If sum exists, it is returned
	//If sum does not exists, return an error (nil)
	sum, exists := ss.sums[key{x, y}]
	if !exists {
		return -1 //Assume -1 means nil
	}
	return
}

func (ss *addService) DeleteSum(x int, y int) {
	delete(ss.sums, key{x, y})
}

// Make the service available on the network (i.e., callable by other services)
// Function name must be 'ServeHTTP'
func (ss *addService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	x, _ := strconv.Atoi(num1)
	y, _ := strconv.Atoi(num2)

	switch r.Method {
	case "POST":
		ss.PostSum(x, y)
		fmt.Fprintf(w, "New sum saved!")
	case "GET":
		sum := ss.GetSum(x, y)
		if sum == -1 {
			fmt.Fprintf(w, "Sum does not exist in cache!")
		} else {
			fmt.Fprintf(w, "Sum = %v", sum)
		}
	case "PUT":
	case "DELETE":
		ss.DeleteSum(x, y)
		fmt.Fprintf(w, "Sum deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}

func main() {
	//Create the cache
	ss := &addService{
		sums: map[key]int{
			key{1, 2}: 3,
		},
	}

	http.ListenAndServe("localhost:9999", ss)
}

// URL
// http://localhost:9999?x=1&y=3
