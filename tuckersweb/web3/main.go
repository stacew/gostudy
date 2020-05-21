package main

import (
	"net/http"

	"stacew/gostudy/tuckersweb/web3/myapp"
)

func main() {
	http.ListenAndServe(":8080", myapp.NewHandler())
}
