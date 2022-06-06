package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	pathToYaml := flag.String("yamlsrc", "", "Path to yaml file with matching between short and long urls.")
	flag.Parse()

	var yaml 
	if pathToYaml != "" {
		file, err := os.Open(pathToYaml)
		if err != nil {
			log.Fatal(err)
		}
		reader := yaml.NewDecoder(file)
		reader.Decode()
		defer file.Close()
	} else {
		var yaml byte = `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// 	json := `
	// 	[
	// 		{
	// 			"path": "/urlshort",
	// 			 "url": "https://github.com/gophercises/urlshort"
	// 		},
	// 		{
	// 			"path": "/urlshort-final",
	// 			 "url": "https://github.com/gophercises/urlshort/tree/solution"
	// 		}
	// ]`
	// 	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
