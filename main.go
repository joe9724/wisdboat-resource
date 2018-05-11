package main

import (
	"net/http"
	"log"
)


func main() {
	http.Handle("/resource/",http.StripPrefix("/resource",http.FileServer(http.Dir("/root/go/src/resource"))))

	err := http.ListenAndServe(":95", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}