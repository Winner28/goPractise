package main

import "muXhttp"
import "net/http"


func main() {
	http.ListenAndServe("localhost:8000", muXhttp.InitAll())
}