package dataModel

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3" //위의 sql에 대한 go-sqlite3를 사용한다는 암시적 의미로 추가
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos WHERE sessionId=?", sessionId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}

	return todos
}
func (s *sqliteHandler) AddTodo(sessionId string, name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (sessionId, name, completed, createdAt) VALUES (?, ?, ?, datetime('now'))")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(sessionId, name, false)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId() //마지막 추가된 아이디 값 나옴.
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}
func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	cnt, _ := rst.RowsAffected()
	return cnt > 0
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}
func (s *sqliteHandler) Close() {
	s.db.Close()
}
func newSqliteHandler(filepath string) DataHandlerInterface {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			sessionId	STRING,
			name		TEXT,
			completed	BOOLEAN,
			createdAt	DATETIME
		);

		CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON todos (
			sessionId ASC
		);`)
	// Index 작업. 기존 DB 순차탐색이라 느리기 떄문에 이렇게 처리하려 map key value 느낌으로
	// 추가 될 때 Index 트리를 만듬.
	statement.Exec()

	return &sqliteHandler{db: database}
}
