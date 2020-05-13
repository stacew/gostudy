package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserClass struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello Root")

	//테스트 용
	pUser := &UserClass{} //힙 메모리 pUser := new(UserClass) 로도 쓸수 있음.
	oUser := UserClass{}  //스택 메모리

	pUser.CreatedAt = time.Now()
	pUser = nil //nil 대입 시, 힙 메모리 레퍼런스 체크 후 해제. delete 예약어 없음.
	//pUser.CreatedAt = time.Now() 크래시

	oUser.CreatedAt = time.Now()
	//oUser = nil //불가능
}
func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") //localhost/bar?name='yslee'
	if name == "" {
		name = "No Name"
	}

	fmt.Fprintf(w, "hello %s!", name)
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pUser := new(UserClass)
	err := json.NewDecoder(r.Body).Decode(pUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request : ", err)
		return
	}

	pUser.CreatedAt = time.Now()

	data, _ := json.Marshal(pUser)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))

	pUser = nil //
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	return mux
}
