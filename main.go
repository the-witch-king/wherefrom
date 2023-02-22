package main

// 146 - Ootsuka, Akio, MP100

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	Data []JikanPersonVoice `json:"data"`
}

type JikanPersonVoice struct {
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

	if len(os.Args) < 3 {
		printHelp()
		os.Exit(420)
	}

	userName := os.Args[1]
	actorId := os.Args[2]

	if userName == "" || actorId == "" {
		os.Exit(69)
	}

	seenAnime, err := getUsersSeenAnime(userName)
	if err != nil {
		log.Fatalf("Unable to retrieve user's seen anime.\nOriginal error: %v", err)
	}

	fmt.Println("For actor with ID: ", actorId)
	voiceRoles, err := getVoiceRoles(actorId)
	if err != nil {
		log.Fatalf("Unable to retrieve actor's voice roles.\nOriginal error: %v", err)
	}

	if len(voiceRoles) < 1 {
		fmt.Println("You haven't seen anything that this person has voice acted in.")
		os.Exit(42)
	}

	fmt.Printf("\nYou've seen them in: \n\n=============================\n")
	tbl := table.New("Show", "Character")

	for _, vr := range voiceRoles {
		if anime, seen := seenAnime[fmt.Sprintf("%d", vr.Anime.MalID)]; seen {
			// fmt.Printf("\n[%s]:\t\t[%s]!", anime.Node.Title, voiceRole.Character.Name)
			tbl.AddRow(anime.Node.Title, vr.Character.Name)
		}
	}
	tbl.Print()
}

func getUsersSeenAnime(userName string) (map[string]MALAnime, error) {
	seenAnime := map[string]MALAnime{}
	nextUrl := fmt.Sprintf("https://api.myanimelist.net/v2/users/%s/animelist", userName)

	for nextUrl != "" {
		req, err := http.NewRequest("GET", nextUrl, nil)

		if err != nil {
			return nil, err
		}

		req.Header.Add("X-MAL-CLIENT-ID", os.Getenv(CLIENT_ID_KEY))

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if isHttpError(resp) {
			return nil, generateHttpError(resp)
		}

		responseData := MALUserAnimeListResponse{}
		err = json.NewDecoder(resp.Body).Decode(&responseData)

		if err != nil {
			return nil, err
		}

		for _, anime := range responseData.Data {
			seenAnime[fmt.Sprintf("%d", anime.Node.Id)] = anime
		}

		nextUrl = responseData.Paging.Next
	}

	return seenAnime, nil
}

func getVoiceRoles(actorId string) ([]JikanPersonVoice, error) {
	personUrl := fmt.Sprintf("https://api.jikan.moe/v4/people/%s/voices", actorId)
	resp, err := http.Get(personUrl)

	if err != nil {
		return nil, err
	}

	if isHttpError(resp) {
		return nil, generateHttpError(resp)
	}

	voiceRoles := JikanGetPersonVoicesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&voiceRoles)

	if err != nil {
		return nil, err
	}

	if len(voiceRoles.Data) < 1 {
		return []JikanPersonVoice{}, nil
	}

	return voiceRoles.Data, nil
}

func isHttpError(resp *http.Response) bool {
	return resp.StatusCode < 200 || resp.StatusCode > 299
}

func generateHttpError(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errorBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			errorBody = []byte("Unable to parse response body")
		}

		return fmt.Errorf("Server error.\nStatus Code: %d\nResponse: %v", resp.StatusCode, string(errorBody))
	}

	return nil
}
