package apitests

import "fmt"
import "net/http"
import "testing"

var initCatId string

func init() {
	// Preparation: delete all existing & create a cat
	ids := []string{}
	call("GET", "/cats", nil, nil, &ids)

	for _, id := range ids {
		call("DELETE", "/cats/" + id, nil, nil, nil)
	} 

	// Create a single cat into the DB
	call("POST", "/cats", &CatModel{Name: "Toto"}, nil, &initCatId)
}

func TestGetCats(t *testing.T) {

	code := 0
	result := []string{}
	err := call("GET", "/cats", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats ->", code, result)

	if code != http.StatusOK {
		t.Error("We should get code 200, got", code)
	}

	if len(result) != 2 {
		t.Error("We should get one item, got", len(result))
		return
	}

	if result[1] != initCatId {
		t.Error("Listing the IDs, got", result[0])
	}
}

// Continue implementing here ...
