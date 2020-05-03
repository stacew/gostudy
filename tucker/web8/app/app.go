package app

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stacew/gostudy/tucker/web8/dataModel"

	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type Success struct {
	Success bool `json:"success"`
}

type AppHandler struct {
	http.Handler //embeded is-a같은 has-a 관계라는데, 이름 정해주면 안 됨...
	dmHandler    dataModel.DataHandlerInterface
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.dmHandler.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.dmHandler.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ok := a.dmHandler.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.dmHandler.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) Close() {
	a.dmHandler.Close()
}

func MakeNewHandler(filepath string) *AppHandler {

	r := mux.NewRouter()

	a := &AppHandler{
		Handler:   r,
		dmHandler: dataModel.NewDataHandler(filepath),
	}

	r.HandleFunc("/todoH", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todoH", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todoH/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todoH/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/", a.indexHandler)

	return a
}
