package t_insert_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestInsertInventory(t *testing.T) {
	//data := map[string]any{"field1": "value1", "field2": "value2"}

	data := []byte(`{"product_id":312, "quantity":4}`)

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:8089/addinventory", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println("response", string(body))

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
