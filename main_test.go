package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/rtravitz/culture_knights"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(os.Getenv("CULTURE_DB_TEST"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"Sullivan Julius"}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "Sullivan Julius" {
		t.Errorf("Expected user name to be 'Sullivan Julius'. Got '%v'", m["name"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{"name":"Cassie Boo"}`)

	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
	}

	if m["name"] == originalUser["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["name"], m["name"], m["name"])
	}

}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO users(name) VALUES($1)", "User "+strconv.Itoa(i))
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
	id SERIAL,
	name TEXT NOT NULL
)`
