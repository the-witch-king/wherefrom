package main

// 146 - Ootsuka, Akio, MP100

import (
	"log"
	"net/http"
	"os"


	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const CLIENT_ID_KEY = "MAL_CLIENT_ID"

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var MalClientId = os.Getenv(CLIENT_ID_KEY)

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string("foo"),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(show)
}
