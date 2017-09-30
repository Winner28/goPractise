package main

import (
	"api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started on port 7000")
	http.ListenAndServe(":7000", api.Handlers())
}
