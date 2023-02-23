package main  

import (
	"fmt"
	"log"
	"os"

	"wherefrom/m/v2/jikan"
	"wherefrom/m/v2/mal"
	"wherefrom/m/v2/shared"
)

const CLIENT_ID_KEY = "MAL_CLIENT_ID"
var MAL_CLIENT_ID = os.Getenv(CLIENT_ID_KEY)

func getActorInSeenAnimes(malUserName string, actorId string) []shared.Appearance {
	mc := mal.MakeMALClient(MAL_CLIENT_ID)
	seenAnimes, err := mc.GetUserSeenAnime(malUserName)

	if err != nil {
		log.Fatalf("Unable to retrieve seen animes: %v", err)
	}

	jc := jikan.MakeJikanClient()
	voiceRoles, err := jc.GetPersonFull(actorId)

	if err != nil {
		log.Fatalf("Unable to retrieve voice actor roles: %v", err)
	}

	seenIn := []shared.Appearance{}

	for _, vr := range voiceRoles {
		if anime, seen := seenAnimes[fmt.Sprintf("%d", vr.Anime.MalID)]; seen {
			seenIn = append(seenIn, shared.Appearance{Show: anime.Node.Title, Character: vr.Character.Name})
		}
	}

	return seenIn
}

func main() {
	userName := os.Args[1]
	actorId := os.Args[2]

	appearances := getActorInSeenAnimes(userName, actorId)

	for _, a := range appearances {
		fmt.Printf("\n%s: %s", a.Show, a.Character)
	}
}
