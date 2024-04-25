package main

import (
	"log"
	"net/http"
)

// this version must follow the same versioning scheme as the tag
const latestVersion = "v0.1.1"

func main() {
	versionHandlers := map[string]handlerFunc{
		latestVersion: handleV2,
		"v0.1.0":      handleV1,
	}

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		version := r.Header.Get("X-Api-Version")
		handler, exists := versionHandlers[version]

		if !exists {
			handler = versionHandlers[latestVersion]
		}

		if err := handler(w, r); err != nil {
			http.Error(w, "Failed to generate payload", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
