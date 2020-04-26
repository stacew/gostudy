package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (user User) IsOld() bool {
	return user.Age > 30
}

func main() {
	user := User{Name: "ys", Email: "aaa@naver.com", Age: 11}
	user2 := User{Name: "ys2", Email: "bbb@naver.com", Age: 55}
	users := []User{user, user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl",
		"templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	//tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	//tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
	//{{range .}}으로 변경
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
