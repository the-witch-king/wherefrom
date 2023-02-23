package mal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wherefrom/m/v2/utils"
)

type MalClient interface {
	GetUserSeenAnime(userName string) (map[string]MALAnime, error)
}

type client struct {
	malClientId string
}

func MakeMALClient(malClientId string) MalClient {
	return &client{malClientId}
}

func (c client) GetUserSeenAnime(userName string) (map[string]MALAnime, error) {
	seenAnime := map[string]MALAnime{}
	nextUrl := fmt.Sprintf("https://api.myanimelist.net/v2/users/%s/animelist", userName)

	for nextUrl != "" {
		req, err := http.NewRequest("GET", nextUrl, nil)

		if err != nil {
			return nil, err
		}

		req.Header.Add("X-MAL-CLIENT-ID", c.malClientId)

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if utils.IsHttpError(resp) {
			return nil, utils.GenerateHttpError(resp)
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
