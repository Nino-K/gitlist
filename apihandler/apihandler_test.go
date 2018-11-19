package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler func(w http.ResponseWriter, r *http.Request)

func TestGetEmptyArgs(t *testing.T) {
	t.Log("Returns an error when empty url passed in")
	{
		_, err := get("")
		if err == nil {
			t.Error("Expect an error")
		}
		if _, ok := err.(*apiHandlerError); !ok {
			t.Error("Expected apiHandlerError type")
		}

	}
}

func TestGetResponse(t *testing.T) {
	t.Log("Returns http response when request successful")
	{
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, http.StatusOK)
		}
		ts := httptest.NewServer(http.HandlerFunc(testHandler))
		resp, err := get(ts.URL)
		if err != nil {
			t.Errorf("get: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error("expected 200 status")
		}
		ts.Close()
	}
	t.Log("Returns nil response when request fails")
	{
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something bad", http.StatusBadRequest)
		}
		ts := httptest.NewServer(http.HandlerFunc(testHandler))
		resp, err := get(ts.URL)
		if err == nil {
			t.Error("Expected an err")
		}
		if resp != nil {
			t.Error("Expected nil response")
		}
		ts.Close()
	}
}

func TestGetHeader(t *testing.T) {
	t.Log("Should add required header to request")
	{
		var requiredHeader string
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			requiredHeader = r.Header.Get("Accept")
			fmt.Fprintln(w, http.StatusOK)
		}
		ts := httptest.NewServer(http.HandlerFunc(testHandler))
		_, err := get(ts.URL)
		if err != nil {
			t.Errorf("get: %v", err)
		}
		if requiredHeader != STANDARD_HEADER {
			t.Errorf("Expected the following header: %v", STANDARD_HEADER)
		}
	}
}

func TestGetReposError(t *testing.T) {
	t.Log("Returns an error if request fails")
	{
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something bad happened", http.StatusBadRequest)
		}
		ts := httptest.NewServer(http.HandlerFunc(testHandler))
		apiHandler := New(ts.URL)
		repos, err := apiHandler.GetRepos()
		if err == nil {
			t.Error("Expected an error from GetRepos")
		}
		if repos != nil {
			t.Error("Expected nil repositories")
		}
	}
}

func TestGetRepos(t *testing.T) {
	t.Log("Returns a list of repository when request successful")
	{
		testRepos := GetTestRepos("first_repo", "second_repo", "third_repo")
		responseBody, err := json.Marshal(testRepos)
		if err != nil {
			t.Errorf("Marshalling failed: %v", err)
		}
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(responseBody)
		}
		ts := httptest.NewServer(http.HandlerFunc(testHandler))

		apiHandler := New(ts.URL)
		repos, err := apiHandler.GetRepos()
		if err != nil {
			t.Errorf("Not Expected an error :%v", err)
			return
		}
		VerifyItems(repos, testRepos.Items, t)
	}
}

func TestGetReposUnMarshal(t *testing.T) {
	t.Log("Returns an error if it fails to Unmarshal")
	{
		var jsonBlob = []byte(`[some garbage]`)
		testHandler = func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBlob)
		}

		ts := httptest.NewServer(http.HandlerFunc(testHandler))
		apiHandler := New(ts.URL)
		repos, err := apiHandler.GetRepos()
		if err == nil {
			t.Errorf("Expected an unmarshalling error :%v", err)
		}
		if repos != nil {
			t.Error("Not Expected list of repos")
		}
	}
}
