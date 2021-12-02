package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var wg = sync.WaitGroup{}

func main()  {
	r := mux.NewRouter()
	
	// defer wg.Wait()
	
	r.HandleFunc("/{number}", serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	
	// wg.Add(1)
	
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	// params["number"]
	num, _ := strconv.Atoi(params["number"])

	for i := 2; i < num; i++ {
		if num % i == 0 {
			json.NewEncoder(w).Encode(i)
			// wg.Done()
			return
		}
	}

	json.NewEncoder(w).Encode("number is prime")
	// wg.Done()

}
