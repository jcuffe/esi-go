package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jcuffe/esi-go/oauth"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(ctx context.Context, request Request) (Response, error) {
	params := request.QueryStringParameters
	code := params["code"]
	
	conf := oauth.Config
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}
	
	body, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}
	
	// Respond with user data
	response := Response{StatusCode: 200, Body: string(body)}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
