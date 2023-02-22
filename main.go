package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const CLIENT_ID_KEY = "MAL_CLIENT_ID"

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var MalClientId = os.Getenv(CLIENT_ID_KEY)

func getActorInSeenAnimes(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userPayload := &UserPayload{}
	err := json.Unmarshal([]byte(req.Body), userPayload)

	if err != nil {
		return errorHandler(err)
	}

	seenAnimes, err := getUsersSeenAnime(userPayload.UserName)

	if err != nil {
		return errorHandler(err)
	}

	voiceRoles, err := getVoiceRoles(userPayload.ActorId)

	if err != nil {
		return errorHandler(err)
	}

	seenIn := []Appearance{}

	for _, vr := range voiceRoles {
		if anime, seen := seenAnimes[fmt.Sprintf("%d", vr.Anime.MalID)]; seen {
			seenIn = append(seenIn, Appearance{Show: anime.Node.Title, Character: vr.Character.Name})
		}
	}

	responseBody, err := json.Marshal(seenIn)

	if err != nil {
		errorHandler(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}

func errorHandler(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func main() {
	lambda.Start(getActorInSeenAnimes)
}
