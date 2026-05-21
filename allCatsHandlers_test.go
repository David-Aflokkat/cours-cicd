package main

import (
	"net/http"
	"slices"
	"strings"
	"testing"
)

func TestListMapKeys(t *testing.T) {
	input := map[string]Cat{
		"11111111-1111-1111-1111-111111111111": {
			ID:   "11111111-1111-1111-1111-111111111111",
			Name: "Toto",
		},
		"22222222-2222-2222-2222-222222222222": {
			ID:   "22222222-2222-2222-2222-222222222222",
			Name: "Mimi",
		},
	}

	result := listMapKeys(input)

	if len(result) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(result))
	}

	if !slices.Contains(result, "11111111-1111-1111-1111-111111111111") {
		t.Errorf("expected result to contain first cat ID")
	}

	if !slices.Contains(result, "22222222-2222-2222-2222-222222222222") {
		t.Errorf("expected result to contain second cat ID")
	}
}

func TestListCats(t *testing.T) {
	catID := "11111111-1111-1111-1111-111111111111"

	catsDatabase = map[string]Cat{
		catID: {
			ID:        catID,
			Name:      "Toto",
			Color:     "Grey",
			BirthDate: "2023-04-16",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/api/cats", nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	status, body := listCats(req)

	if status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, status)
	}

	ids, ok := body.([]string)
	if !ok {
		t.Fatalf("expected body to be []string, got %T", body)
	}

	if len(ids) != 1 {
		t.Fatalf("expected 1 cat ID, got %d", len(ids))
	}

	if ids[0] != catID {
		t.Errorf("expected cat ID %s, got %s", catID, ids[0])
	}
}

func TestCreateCatValidJSON(t *testing.T) {
	catsDatabase = map[string]Cat{}

	body := strings.NewReader(`{
		"name": "Mimi",
		"color": "Black",
		"birthDate": "2024-01-01"
	}`)

	req, err := http.NewRequest(http.MethodPost, "/api/cats", body)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	status, responseBody := createCat(req)

	if status != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, status)
	}

	newCatID, ok := responseBody.(string)
	if !ok {
		t.Fatalf("expected response body to be string, got %T", responseBody)
	}

	if newCatID == "" {
		t.Fatalf("expected generated cat ID, got empty string")
	}

	createdCat, exists := catsDatabase[newCatID]
	if !exists {
		t.Fatalf("expected cat to be saved in database")
	}

	if createdCat.ID != newCatID {
		t.Errorf("expected cat ID %s, got %s", newCatID, createdCat.ID)
	}

	if createdCat.Name != "Mimi" {
		t.Errorf("expected name Mimi, got %s", createdCat.Name)
	}

	if createdCat.Color != "Black" {
		t.Errorf("expected color Black, got %s", createdCat.Color)
	}

	if createdCat.BirthDate != "2024-01-01" {
		t.Errorf("expected birthDate 2024-01-01, got %s", createdCat.BirthDate)
	}
}

func TestCreateCatInvalidJSON(t *testing.T) {
	catsDatabase = map[string]Cat{}

	body := strings.NewReader(`{ invalid json }`)

	req, err := http.NewRequest(http.MethodPost, "/api/cats", body)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	status, responseBody := createCat(req)

	if status != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if responseBody != "Invalid JSON input" {
		t.Errorf("expected invalid JSON error message, got %v", responseBody)
	}

	if len(catsDatabase) != 0 {
		t.Errorf("expected database to remain empty, got %d entries", len(catsDatabase))
	}
}