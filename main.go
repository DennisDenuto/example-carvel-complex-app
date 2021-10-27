package main

import (
	"net/http"
)

func main() {
	println("Listening on :8080")

	err := http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(`Hello Carvel From Complex App!`))
		if err != nil {
			panic(err.Error())
		}
	}))
	if err != nil {
		panic(err.Error())
	}
}
