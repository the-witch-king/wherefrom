package main

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	"wherefrom/m/v2/mal"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

const LAMBDA_ENV = "lambda"
const DEV_ENV = "dev"

//go:embed frontend/build
var build embed.FS

var malClient = mal.MakeMALClient(os.Getenv("MAL_CLIENT_ID"))

func getPage() http.Handler {
	fsys := fs.FS(build)
	html, _ := fs.Sub(fsys, "frontend/build")

	return http.FileServer(http.FS(html))
}

func getUserAnimeList(w http.ResponseWriter, r *http.Request) {
	p := UserListRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Unable to parse request.")
		return
	}

	malUserName := p.UserName
	userSeenAnime, err := malClient.GetUserSeenAnime(malUserName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Unable to get user's anime list")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userSeenAnime)
}

func unfoundRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Sorry, only GET and POST methods are supported.")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func router(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

	env, found := os.LookupEnv("ENV")

	if !found {
		log.Fatalf("No environment set.")
	}

	if env == LAMBDA_ENV {
		lambda.Start(httpadapter.New(mux).ProxyWithContext)
	} else {
		err := http.ListenAndServe(":3333", mux)
		if err != nil {
			log.Fatalf("Unable to start server.")
		}
	}
}
