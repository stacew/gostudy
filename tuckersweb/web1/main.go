package main

import (
	"net/http"

	"github.com/stacew/gostudy/tuckersweb/web1/myapp"
)

func main() {
	http.ListenAndServe(":8080", myapp.NewHttpHandler())
}
