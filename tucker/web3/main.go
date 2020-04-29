package main

import (
	"net/http"

	"github.com/stacew/gostudy/tucker/web3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
