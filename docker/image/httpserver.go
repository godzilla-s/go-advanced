package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> hello :%v<h1>", time.Now())
}

func imageList(w http.ResponseWriter, r *http.Request) {

}

func containerList(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Println("http server start...")
	http.HandleFunc("/", handle)
	http.HandleFunc("/images", imageList)
	http.HandleFunc("/containers", containerList)
	http.ListenAndServe(":8010", nil)
}
