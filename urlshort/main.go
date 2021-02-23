package main

import (
	"fmt"
	"net/http"

	"github.com/marcosgeo/urlshort/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the Maphandler using the mux as the fallback
	pathsToUrl := map[string]string{
		"/urlshor-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":    "https://godoc.org/gopgk.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrl, mux)

	// Build the YAMLHnalder using the mapHandler as the fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, world!")
}
