package main

import "net/http"

import "github.com/gorilla/mux"

func getCat(req *http.Request) (int, any) {
	catID := mux.Vars(req)["catId"]
	Logger.Info("Getting the cat: ", catID)

	if cat, found := catsDatabase[catID]; found {
		Logger.Info("Cat found")
		return http.StatusOK, cat
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}
