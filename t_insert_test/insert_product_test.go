package t_insert_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertProduct(t *testing.T) {
	var response string

	data := []byte(`{"product_id":230,"name":"rocking chair", "specification":{"made":"china"}, "sku":"i91u0", "category_id":3, "price":12}`)

	resp, err := http.Post("http://localhost:8089/addproduct", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if response != "product has been inserted" {
		t.Errorf("Expected product has been inserted, got %s", response)
	}

}
