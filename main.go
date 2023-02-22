package main

// 146 - Ootsuka, Akio, MP100

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rodaine/table"
)

const CLIENT_ID_KEY = "MAL_CLIENT_ID"

type MALUserAnimeListResponse struct {
	Data   []MALAnime `json:"data"`
	Paging MALPaging  `json:"paging"`
}

type MALAnime struct {
	Node struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		MainPicture struct {
			Medium string `json:"medium"`
			Large  string `json:"large"`
		}
	} `json:"node"`
	ListStatus struct {
		Status             string `json:"status"`
		Score              int    `json:"score"`
		NumWatchedEpisodes int    `json:"num_watched_episodes"`
		IsRewatching       bool   `json:"is_rewatching"`
		UpdatedAt          string `json:"updated_at"`
	} `json:"list_status"`
}

type MALPaging struct {
	Next string `json:"next"`
}

type JikanGetPersonVoicesResponse struct {
	Data []JikanPersonVoices `json:"data"`
}

type JikanPersonVoices struct {
	Role  string `json:"role"`
	Anime struct {
		MalID  int    `json:"mal_id"`
		URL    string `json:"url"`
		Images struct {
			Jpg struct {
			} `json:"jpg"`
			Webp struct {
			} `json:"webp"`
		} `json:"images"`
		Title string `json:"title"`
	} `json:"anime"`
	Character struct {
		MalID  int    `json:"mal_id"`
		URL    string `json:"url"`
		Images struct {
		} `json:"images"`
		Name string `json:"name"`
	} `json:"character"`
}

func printHelp() {
	fmt.Println("USAGE:")
	fmt.Println("<command to run> <user name from MAL> <actor id>")
	fmt.Println("ie: ./main some-random-guy 146 # voice actor for Mob")
}

func main() {
	seenAnime := map[string]MALAnime{}

	if len(os.Args) < 3 {
		printHelp()
		os.Exit(420)
	}

	userWithList := os.Args[1]
	actorId := os.Args[2]

	if userWithList == "" || actorId == "" {
		os.Exit(69)
	}

	nextUrl := fmt.Sprintf("https://api.myanimelist.net/v2/users/%s/animelist", userWithList)
	for nextUrl != "" {
		req, err := http.NewRequest("GET", nextUrl, nil)

		if err != nil {
			panic("FUCK")
		}
		req.Header.Add("X-MAL-CLIENT-ID", os.Getenv(CLIENT_ID_KEY))

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			panic("FUIDGE")
		}

		defer resp.Body.Close()

		responseData := MALUserAnimeListResponse{}
		err = json.NewDecoder(resp.Body).Decode(&responseData)

		if err != nil {
			panic(err)
		}

		nextUrl = responseData.Paging.Next
		for _, anime := range responseData.Data {
			seenAnime[fmt.Sprintf("%d", anime.Node.Id)] = anime
		}
	}

	fmt.Println("For actor with ID: ", actorId)

	// Get animes from person
	personUrl := fmt.Sprintf("https://api.jikan.moe/v4/people/%s/voices", actorId)
	personResp, err := http.Get(personUrl)

	if err != nil {

		panic("NONONO")
	}

	voiceRoles := JikanGetPersonVoicesResponse{}
	err = json.NewDecoder(personResp.Body).Decode(&voiceRoles)

	if len(voiceRoles.Data) < 1 {
		fmt.Println("You haven't seen anything that this person has voice acted in.")
		os.Exit(42)
	}

	fmt.Printf("\nYou've seen them in: \n\n=============================\n")
	tbl := table.New("Show", "Character")

	for _, voiceRole := range voiceRoles.Data {
		if anime, seen := seenAnime[fmt.Sprintf("%d", voiceRole.Anime.MalID)]; seen {
			// fmt.Printf("\n[%s]:\t\t[%s]!", anime.Node.Title, voiceRole.Character.Name)
			tbl.AddRow(anime.Node.Title, voiceRole.Character.Name)
		}
	}

	tbl.Print()
}
