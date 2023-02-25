package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

//go:embed static
var mainPage embed.FS

func getPage() http.Handler {
	fsys := fs.FS(mainPage)
	html, _ := fs.Sub(fsys, "static")

	return http.FileServer(http.FS(html))
}

func getUserAnimeList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{ \"name\": \"Foo!\" }")
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPage().ServeHTTP(w, r)
	case "POST":
		getUserAnimeList(w, r)
	default:
		getPage().ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	// mux.Handle("/", router)
	// mux.HandleFunc("/", router)
	mux.HandleFunc("/", router)

	err := http.ListenAndServe(":3333", mux)

	if err != nil {
		fmt.Printf("\nUnable to start server.\n")
		os.Exit(69)
	}
}
