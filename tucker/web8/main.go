package main

import (
	"net/http"

	"github.com/stacew/gostudy/tucker/web8/app"
	"github.com/urfave/negroni"
)

func main() {

	mux := app.MakeNewHandler("./test.db") //실행인자 이용 가능 flag.Arg
	defer mux.Close()

	neg := negroni.Classic()
	neg.UseHandler(mux)

	http.ListenAndServe(":3000", neg)
}
