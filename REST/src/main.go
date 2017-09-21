package main

import (
	"api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":7000", api.Handlers())
}
