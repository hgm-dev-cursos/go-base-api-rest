package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpRouter := http.NewServeMux()

	httpRouter.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Printf("Method being called -> " + request.Method)
		writer.WriteHeader(201)
		_, _ = writer.Write([]byte("Hello World!"))
	})

	err := http.ListenAndServe(":8000", httpRouter)
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
