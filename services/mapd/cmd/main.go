package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	mapd "github.com/nk-designz/metroDB/services/mapd/client"
)

var (
	mapdInstance *mapd.Mapd
)

func serve() {
	log.Println("started serving on port:", os.Args[3])
	router := mux.NewRouter()
	router.HandleFunc("/{key}", GetServeEndpoint).Methods("GET")
	router.HandleFunc("/{key}", SetServeEndpoint).Methods("POST")
	router.HandleFunc("/{key}", SetServeEndpoint).Methods("PUT")
	router.HandleFunc("/{key}", DeleteServeEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Args[3]), router))
}

func GetServeEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("received GET on", params["key"])
	value, err := mapdInstance.Get(params["key"])
	if err == nil {
		w.Write(value)
	} else {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
}

func SetServeEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	body, err := ioutil.ReadAll(req.Body)
	log.Println("received SET on", params["key"])
	_, err = mapdInstance.Set(params["key"], body)
	if err == nil {
		w.Write([]byte(params["key"]))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
}

func DeleteServeEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("received DELETE on", params["key"])
	_, err := mapdInstance.Set(params["key"], []byte{})
	if err == nil {
		w.Write([]byte(params["key"]))
	} else {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
}

func main() {
	address := os.Args[1]
	mapdInstance = mapd.New(address)
	if os.Args[2] == "setsafe" {
		mapdInstance.SetSafe(
			os.Args[3],
			[]byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "set" {
		mapdInstance.Set(
			os.Args[3],
			[]byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "setfromfile" {
		fileContent, err := ioutil.ReadFile(os.Args[4])
		if err != nil {
			fmt.Println("Could not open file")
			panic(err)
		}
		mapdInstance.Set(
			os.Args[3],
			fileContent)
		fmt.Println("Seems ok")
	} else if os.Args[2] == "get" {
		byteval, _ := mapdInstance.Get(os.Args[3])
		fmt.Println(fmt.Sprintf("%s", byteval))
	} else if os.Args[2] == "serve" {
		serve()
	} else {
		fmt.Println("No valid command")
	}
}
