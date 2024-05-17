package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Home Page</title>
	</head>
	<body>
		<h1>Hello, world from index routes in Go.</h1>
		<p><a href="http://localhost:3000/">Go to localhost:3000</a></p>
	</body>
	</html>`
	w.Write([]byte(html))
}
