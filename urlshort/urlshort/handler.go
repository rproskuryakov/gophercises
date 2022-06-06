package urlshort

import (
	encodingJson "encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

type urlConfig struct {
	url  string `yaml:"url" json:"url"`
	path string `yaml:"path" json:"path"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var config []urlConfig
	err := yaml.Unmarshal(yml, &config)
	if err != nil {
		return nil, err
	}
	urlMap := map[string]string{}
	for _, value := range config {
		urlMap[value.path] = value.url
	}
	return MapHandler(urlMap, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var config []urlConfig
	err := encodingJson.Unmarshal(json, &config)
	if err != nil {
		return nil, err
	}
	urlMap := map[string]string{}
	for _, value := range config {
		urlMap[value.path] = value.url
	}
	return MapHandler(urlMap, fallback), nil
}
