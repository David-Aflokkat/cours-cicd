package main

import "net/http"


func getCat(req *http.Request) (int, any) {
	Logger.Info("Getting a cat")

	catID := req.PathValue("catId")

	cat, exists := catsDatabase[catID]
	if !exists {
		return http.StatusNotFound, "il n'existe aucun chat pour cet id"
	}

	return http.StatusOK, cat
}

func deleteCat(req *http.Request) (int, any) {
	Logger.Info("Deleting a cat")

	catID := req.PathValue("catId")

	_, exists := catsDatabase[catID]
	if !exists {
		return http.StatusNotFound, "il n'existe aucun chat pour cet id"
	}

	delete(catsDatabase, catID)

	return http.StatusNoContent, nil
}