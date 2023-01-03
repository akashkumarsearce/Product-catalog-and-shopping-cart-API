package t_get_test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetInventory(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/inventorydetail/4")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
