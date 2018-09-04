package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name: "Index", Method: "GET", Pattern: "/", HandleFunc: Index},
	Route{Name: "Time", Method: "GET", Pattern: "/time", HandleFunc: Time},
}

type Handler struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Handlers []Handler

// 接口实现
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To Index")
}

func Time(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Now: ", time.Now())
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("CofoxAPI: %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

func main() {
	//http.HandleFunc("/", handle)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
