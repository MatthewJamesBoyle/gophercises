package main

import (
	"fmt"
	"net/http"
	"gophercises/urlshortner"
	"flag"
	"io/ioutil"
)

func main() {
	mux := defaultMux()

	yamlFile := flag.String("yml", "default.yml", "yml file to parse, defaults to default.yml ")

	flag.Parse()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yaml,err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		panic(err)
	}

	fmt.Print(yaml)

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
	fmt.Fprintln(w, "Hello, world!")
}
