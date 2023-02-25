package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

//go:embed frontend/build
var build embed.FS

func getPage() http.Handler {
	fsys := fs.FS(build)
	html, _ := fs.Sub(fsys, "frontend/build")

	return http.FileServer(http.FS(html))
}

func getUserAnimeList(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	io.WriteString(w, "{ \"items\": [\"a\", \"b\", \"c\"] }")
}

func unfoundRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Sorry, only GET and POST methods are supported.")
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPage().ServeHTTP(w, r)
	case "POST":
		getUserAnimeList(w, r)
	default:
		unfoundRoute(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	err := http.ListenAndServe(":3333", mux)

	if err != nil {
		fmt.Printf("\nUnable to start server.\n")
		os.Exit(69)
	}
}
