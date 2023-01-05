package HandlerCategory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category typedefs.Category_master = typedefs.Category_master{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &category)
	fmt.Println(category.Category_Id, category.Category_Name)
	//unmarshal the json values from postman to put into database
	db := DbConnect.ConnectToDB()
	stmt, err := db.Prepare(queries.AddCategory)
	_, err = stmt.Exec(category.Category_Id, category.Category_Name)

	if err != nil {
		fmt.Println(err) //check here
	} else {
		Helpers.SendJResponse(response.CategoryDetailAdded, w)
	}

}
