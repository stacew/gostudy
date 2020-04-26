package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	println("STD print TestIndexPage() - Get Start")
	resp, err := http.Get(ts.URL)
	println("STD print TestIndexPage() - Get End")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))

	println()
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf) //log : std output -> buffer output

	println("STD print TestDecoHandler() - Get Start")
	resp, err := http.Get(ts.URL)
	println("STD print TestDecoHandler() - Get End")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER2] Started")
	//assert.Contains(string(line), "[LOGGER1] Started")
}
