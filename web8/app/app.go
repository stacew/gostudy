package app

import (
	"net/http"
	"time"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

//var todoMap map[int]*Todo
var todoSlice []*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//	list := []*Todo{}
	//	for _, v := range todoMap {
	//		list = append(list, v)
	//	}
	//
	//	rd.JSON(w, http.StatusOK, list)

	rd.JSON(w, http.StatusOK, todoSlice)
}

func addTestTodos() {
	todoSlice = append(todoSlice, &Todo{1, "Buy a milk", false, time.Now()})
	todoSlice = append(todoSlice, &Todo{2, "Excercise", true, time.Now()})
	todoSlice = append(todoSlice, &Todo{3, "Home work", false, time.Now()})
	//	todoMap[1] = &Todo{1, "Buy a milk", false, time.Now()}
	//	todoMap[2] = &Todo{2, "Excercise", true, time.Now()}
	//	todoMap[3] = &Todo{3, "Home work", false, time.Now()}
}

func MakeNewHandler() http.Handler {
	//	todoMap = make(map[int]*Todo)
	addTestTodos()
	rd = render.New()

	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)

	return r
}
