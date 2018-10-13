package main

import (
	"net/http"
	"log"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello, this is version 1!"))
	request.ParseForm()
}

func main() {
	http.HandleFunc("/", hello)
	log.Println("String server ... v1")
	err := http.ListenAndServe(":80", nil)
	if err == nil{
		log.Fatal("Listion and serve: ",err)
	}
}


