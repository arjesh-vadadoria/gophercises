package main

import (
	"fmt"
	"net/http"
	"urlShortner"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlShortner.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /arjesh
  url: https://github.com/arjesh-vadadoria
`
	yamlHandler, err := urlShortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	_ = yamlHandler

	json := `
[
{
"path" : "/arjesh-github",
"url" : "https://github.com/arjesh-vadadoria"
},
{
"path" : "/learn-go",
"url" : "https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors"
}
]
`
	jsonHandler, err := urlShortner.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
