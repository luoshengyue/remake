package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	loggerHandleFunc(r)
	_, err := fmt.Fprintf(w, "Your location:  %s\n", r.URL.Path[1:])
	if err != nil {
		log.Fatalln("index Handle error")
	}
}

func readBody(w http.ResponseWriter, r *http.Request) {
	loggerHandleFunc(r)
	log.Println("Enter readBody handle")
	all, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("readBody - io - ReadAll err")
		return
	}
	_, err = fmt.Fprintf(w, "Body content: %s\n", all)
}

func query(writer http.ResponseWriter, request *http.Request) {
	loggerHandleFunc(request)
	values := request.URL.Query()
	_, _ = fmt.Fprintf(writer, "Query is: %v", values)
	log.Println("Query Success")
}

func form(w http.ResponseWriter, r *http.Request) {
	loggerHandleFunc(r)
	err := r.ParseForm()
	if err != nil {
		log.Fatalln("ParseForm error")
		return
	}
	_, _ = fmt.Fprintf(w, "form content: %v", r.Form)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	loggerHandleFunc(r)
	_, err := fmt.Fprintln(w, "Hello, Gopher")
	if err != nil {
		log.Fatalln("Something wrong with greeting")
	}
}

func loggerHandleFunc(r *http.Request) {
	fmt.Printf("[%s] - %q - %v\n", r.Method, r.URL.String(), time.Now())
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/read", readBody)
	http.HandleFunc("/query", query)
	http.HandleFunc("/form", form)
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(greeting)))
}
