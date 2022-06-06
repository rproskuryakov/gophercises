package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"htmp/template"
	"log"
	"net/http"
	"os"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

type Chapter struct {
	title   string   `json:"title"`
	story   []string `json:"story"`
	options []struct {
		text string `json:"text"`
		arc  string `json:"arc"`
	} `json:"options,omitempty"`
}

func main() {
	filename := flag.String("filename", "gopher.json", "Path to file with adventure data.")
	flag.Parse()
	file, err := os.Open(*filename)
	adventureDescription := make(map[string]Chapter)
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}
	decoder.Decode(&adventureDescription)
	file.Close()

	tmpl, err := template.New("name").Parse("template.html")
	if err != nil {
		log.Fatal(tmpl)
	}

	http.HandleFunc("/", baseHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
