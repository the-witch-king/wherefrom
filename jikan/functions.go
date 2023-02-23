package jikan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wherefrom/m/v2/utils"
)

type JikanClient interface {
	GetVoiceRoles(actorId string) ([]JikanPersonVoice, error)
}

type jikanClient struct{}

func MakeJikanClient() JikanClient {
	return &jikanClient{}
}

func (j jikanClient) GetVoiceRoles(actorId string) ([]JikanPersonVoice, error) {
	personUrl := fmt.Sprintf("https://api.jikan.moe/v4/people/%s/voices", actorId)
	resp, err := http.Get(personUrl)

	if err != nil {
		return nil, err
	}

	if utils.IsHttpError(resp) {
		return nil, utils.GenerateHttpError(resp)
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
