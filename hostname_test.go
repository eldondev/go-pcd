package main

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/brimstone/go-saverequest"
)

func Test_handleHostnameGet(t *testing.T) {
	// need to reset root
	t.Log("Testing hostname")
	req, _ := saverequest.FakeRequest("GET", "/hostname", map[string]string{}, "")
	w := httptest.NewRecorder()
	MyReadFile = func(filename string) ([]byte, error) {
		return []byte("asdf"), nil
	}
	handleHostname(w, req)
	if w.Body.String() != "asdf" {
		t.Errorf("Got unexpected hostname")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got proper hostname")
}

func Test_handleHostnameGetError(t *testing.T) {
	// need to reset root
	t.Log("Testing hostname error")
	req, _ := saverequest.FakeRequest("GET", "/hostname", map[string]string{}, "")
	w := httptest.NewRecorder()
	MyReadFile = func(filename string) ([]byte, error) {
		return []byte(""), fmt.Errorf("This is an error")
	}
	handleHostname(w, req)
	if w.Code != 500 {
		t.Errorf("Got unexpected status code")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got proper status code")
}

func Test_handleHostnamePost(t *testing.T) {
	// need to reset root
	t.Log("Testing hostname POST")
	req, _ := saverequest.FakeRequest("POST", "/hostname", map[string]string{}, "hostname")
	w := httptest.NewRecorder()
	MyWriteFile = func(filename string, contents []byte, mode os.FileMode) error {
		return nil
	}
	handleHostname(w, req)
	if w.Code != 200 {
		t.Errorf("Got unexpected status code")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got proper status code")
}

func Test_handleHostnamePostError(t *testing.T) {
	// need to reset root
	t.Log("Testing hostname POST error")
	req, _ := saverequest.FakeRequest("POST", "/hostname", map[string]string{}, "hostname")
	w := httptest.NewRecorder()
	MyWriteFile = func(filename string, contents []byte, mode os.FileMode) error {
		return fmt.Errorf("This is an error")
	}
	handleHostname(w, req)
	if w.Code != 500 {
		t.Errorf("Got unexpected status code")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got proper status code")
}
