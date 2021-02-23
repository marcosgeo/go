package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler ...
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.Handler, error) {
	// 1. Parse the yaml
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	// 2. Convert YAML array into map
	pathsToUrls := buildMap(pathUrls)

	// 3. return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func parseYaml(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
