package main

import "embed"
import "encoding/json"
import "io/fs"
import "net/http"

import "github.com/gorilla/mux"
import "gitlab.com/ggpack/webstream"


//go:embed swagger-ui
var content embed.FS


func logReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Infof("New request to: '%s %s'", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func newApp() http.Handler {
	Logger.Info("Init the backend")

	router := mux.NewRouter()
	router.HandleFunc("/", getHomeHandler).Methods("GET")
	router.HandleFunc("/api/cats", makeHandlerFunc(createCat)).Methods("POST")
	router.HandleFunc("/api/cats", makeHandlerFunc(listCats)).Methods("GET")
	router.HandleFunc("/api/cats/{catId}", makeHandlerFunc(getCat)).Methods("GET")

	fsys, _ := fs.Sub(content, "swagger-ui")
	router.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", http.FileServer(http.FS(fsys))))

	router.HandleFunc("/ws", wsHandler)
	router.HandleFunc("/logs", webstream.UiHandler("/../ws"))

	return logReq(router)
}

// Simple interface to implement to handle requests
type ServiceFunc func(*http.Request) (int, any)

// Wraps the ServiceFunc to manke a http.HandlerFunc with panic handling and JSON response encoding
func makeHandlerFunc(svcFunc ServiceFunc) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {

		code, body := func(req *http.Request) (code int, body any) {
			// General panic/error handler to keep the server up
			defer func() {
				if r := recover(); r != nil {
					Logger.Error("Recovering from a panic: ", r)
					// Using the named return values
					code = http.StatusInternalServerError
					body = http.StatusText(code)
				}
			}()
			return svcFunc(req)
		}(req)

		// Single response
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(code)
		json.NewEncoder(res).Encode(body)
	}
}
