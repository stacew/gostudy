package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("hello Root", string(data))
}

func TestBarHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("hello No Name!", string(data))
}

func TestBarHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=ys", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("hello ys!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusBadRequest, resp.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"ys","last_name":"lee","email":"ccc@naver.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusCreated, resp.Code)

	user := new(UserClass)
	err := json.NewDecoder(resp.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("ys", user.FirstName)
	assert.Equal("lee", user.LastName)
}
