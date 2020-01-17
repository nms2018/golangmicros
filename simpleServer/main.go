package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeServer(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("welcome home!!"))
}
