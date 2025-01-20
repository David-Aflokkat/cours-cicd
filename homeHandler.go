package main

import "net/http"

func getHomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(`
		<html>
		<title>Cats API</title>
		<link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='0.9em' font-size='80'>😺</text></svg>">
		<style>
		html, body
		{
			width: 100%;
		}
		a
		{
			text-decoration: none;
		}
		</style>
		<body>
			<h2>Software version: ` + version + `</h2>
			<br/>
			<a href="swagger/"><h3>🖥️ Swagger OpenAPI UI</h3></a>
			<br/>
			<a href="logs"><h3>📜 Logs stream</h3></a>
		<body>
		</html>
	`))
}
