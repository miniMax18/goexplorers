package urlshort

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"gopkg.in/yaml.v3"
)

//type UrlMapperEntry struct {
//	Path string `xml:"path" json:"path" yaml:"path"`
//	Url  string `xml:"url" json:"url" yaml:"url"`
//}
type UrlMapperEntry struct {
	Path string `json:"path" yaml:"path"`
	Url  string `json:"url" yaml:"url"`
}
type UrlMapper []UrlMapperEntry

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request url: " + r.URL.String())

		shortenedUrl, exists := pathsToUrls[r.URL.String()]
		if !exists {
			slog.Warn("No url in map")
			fallback.ServeHTTP(w, r)
		} else {
			slog.Info("Redirect...")
			http.Redirect(w, r, shortenedUrl, http.StatusMovedPermanently)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	emptyMap := map[string]string{}
	var urlMapper UrlMapper

	err := yaml.Unmarshal(yml, &urlMapper)
	if err != nil {
		slog.Error("error: " + err.Error())
		return nil, err
	}

	for _, mapper := range urlMapper {
		emptyMap[mapper.Path] = mapper.Url
	}

	return MapHandler(emptyMap, fallback), nil
}

func JSONHandler(jsonInput []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlMap := map[string]string{}
	var urlMapper UrlMapper

	err := json.Unmarshal(jsonInput, &urlMapper)
	if err != nil {
		slog.Error("error: " + err.Error())
		return nil, err
	}

	for _, mapper := range urlMapper {
		urlMap[mapper.Path] = mapper.Url
	}
	return MapHandler(urlMap, fallback), nil
}