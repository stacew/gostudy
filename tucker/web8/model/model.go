package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandlerInterface interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

var dbHandler dbHandlerInterface

func Init() {
	// dbHandler = newMapHandler()
	// dbHandler = newSqliteHandler()
}

func GetTodos() []*Todo {
	return dbHandler.getTodos()
}

func AddTodo(name string) *Todo {
	return dbHandler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return dbHandler.removeTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return dbHandler.completeTodo(id, complete)
}
