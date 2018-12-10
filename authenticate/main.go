package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jcuffe/esi-go/oauth"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {

	// Configure an OAuth client
	urlString := oauth.AuthCodeURL

	// Respond with a redirect
	response := Response{
		StatusCode: 302,
		Headers:    map[string]string{"Location": urlString},
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
