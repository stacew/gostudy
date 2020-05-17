package main

import (
	"net/http"

	"github.com/stacew/gostudy/tuckersweb/web8/app"
)

func main() {

	mux := app.MakeNewHandler("./test.db") //실행인자 이용 가능 flag.Arg
	defer mux.Close()

	http.ListenAndServe(":8080", mux)
}
