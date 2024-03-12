package main_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	// テスト対象のパッケージをインポート
	"/home/crystal9210/Go_util_templates/advance_ex/advance_examples/context/main"
)

func TestContext_Data(t *testing.T) {
	// create a new context
	c := main.NewContext(httptest.NewRecorder(), &http.Request{Method: "GET"})

	// create a response body
	body := []byte("hello world")

	// test the response with a status code of 200 and a content type of "text/plain"
	c.Data(200, "text/plain", body)
	if c.Writer.Status != 200 {
		t.Errorf("Expected status code to be 200, got %d", c.Writer.Status)
	}
	if c.Writer.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Expected content type to be text/plain; charset=utf-8, got %s", c.Writer.Header().Get("Content-Type"))
	}
	data, _ := ioutil.ReadAll(c.Writer.Body)
	if !bytes.Equal(body, data) {
		t.Errorf("Expected response body to be %s, got %s", body, data)
	}

	// test the response with a status code of 404 and a content type of "application/json"
	c.Data(404, "application/json", body)
	if c.Writer.Status != 404 {
		t.Errorf("Expected status code to be 404, got %d", c.Writer.Status)
	}
	if c.Writer.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Expected content type to be application/json; charset=utf-8, got %s", c.Writer.Header().Get("Content-Type"))
	}
	data, _ = ioutil.ReadAll(c.Writer.Body)
	if !bytes.Equal(body, data) {
		t.Errorf("Expected response body to be %s, got %s", body, data)
	}
}
