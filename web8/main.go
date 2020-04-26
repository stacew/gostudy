package main

import (
	"gostudy/web8/app"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {

	mux := app.MakeNewHandler()
	neg := negroni.Classic()

	neg.UseHandler(mux)

	http.ListenAndServe(":3000", neg)
}
