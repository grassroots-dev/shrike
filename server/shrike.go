package main

import (
	"fmt"
	"net/http"

	"github.com/grassroots-dev/shrike/grpc"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", grpc.SayMessage())
	})
	fmt.Println(grpc.SayMessage())
	http.ListenAndServe(":8080", nil)
}
