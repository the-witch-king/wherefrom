package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUsersSeenAnime(userName string) (map[string]MALAnime, error) {
	seenAnime := map[string]MALAnime{}
	nextUrl := fmt.Sprintf("https://api.myanimelist.net/v2/users/%s/animelist", userName)

	for nextUrl != "" {
		req, err := http.NewRequest("GET", nextUrl, nil)

		if err != nil {
			return nil, err
		}

		req.Header.Add("X-MAL-CLIENT-ID", MalClientId)

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
