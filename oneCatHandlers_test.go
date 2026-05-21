package main

import (
	"net/http"
	"testing"
)

func TestGetCatExistingCat(t *testing.T) {
	catID := "11111111-1111-1111-1111-111111111111"

	catsDatabase = map[string]Cat{
		catID: {
			ID:        catID,
			Name:      "Toto",
			Color:     "Grey",
			BirthDate: "2023-04-16",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/api/cats/"+catID, nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	req.SetPathValue("catId", catID)

	status, body := getCat(req)

	if status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, status)
	}

	cat, ok := body.(Cat)
	if !ok {
		t.Fatalf("expected body to be Cat, got %T", body)
	}

	if cat.ID != catID {
		t.Errorf("expected cat ID %s, got %s", catID, cat.ID)
	}

	if cat.Name != "Toto" {
		t.Errorf("expected cat name Toto, got %s", cat.Name)
	}

	if cat.Color != "Grey" {
		t.Errorf("expected cat color Grey, got %s", cat.Color)
	}

	if cat.BirthDate != "2023-04-16" {
		t.Errorf("expected cat birth date 2023-04-16, got %s", cat.BirthDate)
	}
}

func TestGetCatUnknownCat(t *testing.T) {
	catsDatabase = map[string]Cat{}

	catID := "99999999-9999-9999-9999-999999999999"

	req, err := http.NewRequest(http.MethodGet, "/api/cats/"+catID, nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	req.SetPathValue("catId", catID)

	status, body := getCat(req)

	if status != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, status)
	}

	expectedMessage := "il n'existe aucun chat pour cet id"
	if body != expectedMessage {
		t.Errorf("expected message %q, got %v", expectedMessage, body)
	}
}

func TestDeleteCatExistingCat(t *testing.T) {
	catID := "11111111-1111-1111-1111-111111111111"

	catsDatabase = map[string]Cat{
		catID: {
			ID:        catID,
			Name:      "Toto",
			Color:     "Grey",
			BirthDate: "2023-04-16",
		},
	}

	req, err := http.NewRequest(http.MethodDelete, "/api/cats/"+catID, nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	req.SetPathValue("catId", catID)

	status, body := deleteCat(req)

	if status != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, status)
	}

	if body != nil {
		t.Errorf("expected nil body, got %v", body)
	}

	_, exists := catsDatabase[catID]
	if exists {
		t.Errorf("expected cat to be deleted from database")
	}
}

func TestDeleteCatUnknownCat(t *testing.T) {
	catsDatabase = map[string]Cat{}

	catID := "99999999-9999-9999-9999-999999999999"

	req, err := http.NewRequest(http.MethodDelete, "/api/cats/"+catID, nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}

	req.SetPathValue("catId", catID)

	status, body := deleteCat(req)

	if status != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, status)
	}

	expectedMessage := "il n'existe aucun chat pour cet id"
	if body != expectedMessage {
		t.Errorf("expected message %q, got %v", expectedMessage, body)
	}
}