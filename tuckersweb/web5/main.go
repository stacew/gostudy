package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/negroni"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "ys", Email: "ccc@naver.com"}

	//render 사용
	//w.Header().Add("Content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		//render
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	user.CreatedAt = time.Now()
	//render
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
	rd.JSON(w, http.StatusOK, user)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//render 사용
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	//render
	// 	// w.WriteHeader(http.StatusInternalServerError)
	// 	// fmt.Fprint(w, err)
	// 	rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "yslee")
	user := User{Name: "ys", Email: "ccc@naver.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New()
	rd = render.New(render.Options{
		Directory:  "template",                 //default:templates
		Extensions: []string{".html", ".tmpl"}, //default:tmpl
		Layout:     "hello",
	})

	//고릴라/mux -> 고릴라/pat으로 변경
	// mux := mux.NewRouter()
	// mux.HandleFunc("/users", getUserInfoHandler).Methods("GET")
	// mux.HandleFunc("/users", addUserHandler).Methods("POST")
	mux := pat.New()
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	//negroni
	//mux.Handle("/", http.FileServer(http.Dir("public")))
	//http.ListenAndServe(":3000", mux)
	n := negroni.Classic() //public 파일들을 기본적으로 제공
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)

}
