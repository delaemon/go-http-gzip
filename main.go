package main

import (
	"io"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func main() {
	doc := `
	<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title></title>
	<meta charset="utf-8">
	<meta name="description" content="">
	<meta name="author" content="">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="">
	<!--&#91;if lt IE 9&#93;>
	<script src="//cdn.jsdelivr.net/html5shiv/3.7.2/html5shiv.min.js"></script>
	<script src="//cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js"></script>
	<!&#91;endif&#93;-->
	<link rel="shortcut icon" href="">
	</head>
	<body>
	<!-- Place your content here -->
	<!-- SCRIPTS -->
	<!-- Example: <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script> -->
	<h1>Hello</h1>
	</body>
	</html>
	`

	data := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, doc)
	})

	handler := gziphandler.GzipHandler(data)
	//handler := data

	http.Handle("/", handler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
