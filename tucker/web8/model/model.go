package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo
var genID int

//var todoSlice []*Todo

func Init() {
	genID = 0
	todoMap = make(map[int]*Todo)
}

func GetTodo() []*Todo {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	return list
}

func AddTodo(name string) *Todo {
	id := genID //len(todoMap) + 1
	genID++
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	return todo
}

func RemoveTodo(id int) bool {
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		return true
	}
	return false
}

func CompleteTodo(id int, complete bool) bool {

	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}
