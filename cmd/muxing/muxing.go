package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", sayHello)
	router.HandleFunc("/bad", badRequest)
	router.HandleFunc("/data", returnBodyMessage).Methods("POST")
	router.HandleFunc("/headers", returnHeaders).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["PARAM"]

	json.NewEncoder(w).Encode("Hello, " + name + "!")

}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

type Holder struct {
	Param string `json: "PARAM"`
}

func returnBodyMessage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var holder Holder
	json.Unmarshal(reqBody, &holder)

	json.NewEncoder(w).Encode(`I got message: \n` + holder.Param)
}

func returnHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header["A"]
	b := r.Header["B"]
	log.Println(fmt.Printf("json %s %s:\n", a, b))
	//var holder Holder
	//json.Unmarshal(reqBody, &holder)
	//
	aInt, err := strconv.Atoi(a[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	bInt, err := strconv.Atoi(b[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	str := strconv.Itoa(aInt + bInt)
	w.Header().Add("a+b", str)
	log.Println(fmt.Printf("fin %ss:\n", str))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
