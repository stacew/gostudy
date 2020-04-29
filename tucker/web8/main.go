package main

import (
	"net/http"

	"github.com/stacew/gostudy/tucker/web8/app"
	"github.com/urfave/negroni"
)

func main() {

	mux := app.MakeNewHandler()
	neg := negroni.Classic()

	neg.UseHandler(mux)

	http.ListenAndServe(":3000", neg)
}
