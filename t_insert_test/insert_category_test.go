package t_insert_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertCategory(t *testing.T) {
	data := []byte(`{"category_id":12, "name":"home appliances"}`)

	resp, err := http.Post("http://localhost:8089/addcategory", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	var response string
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if response != "category detail has been added successfully" {
		t.Errorf("category detail has been added successfully, got %s", response)
	}

}
