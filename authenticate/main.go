package main

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler() (Response, error) {

	// Generate query string
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("code_challenge_method", "S256")
	q.Set("scope", "publicData esi-skills.read_skills.v1 esi-universe.read_structures.v1")
	q.Set("redirect_uri", "farts")
	q.Set("client_id", "8d404f907fb04aeda0d36d236face0828")
	q.Set("state", "the absolute")
	queryString := q.Encode()

	// Generate EVE authURL for user login
	authURL := url.URL{
		Scheme: "https",
		Host: "login.eveonline.com",
		Path: "/v2/oauth/authorize",
		RawQuery: queryString,
	}

	// Respond with a redirect
	response := Response{
		StatusCode: 302,
		Headers: map[string]string{"Location": authURL.String()},
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
