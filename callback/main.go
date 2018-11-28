package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(ctx context.Context, request Request) (Response, error) {
	params := request.QueryStringParameters
	client := &http.Client{ Timeout: 10 * time.Second }

	tokenURL := "https://login.eveonline.com/v2/oauth/token"

	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", params["code"])

	formData := strings.NewReader(v.Encode())

	req, err := http.NewRequest("POST", tokenURL, formData)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	username := "8d404f907fb04aeda0d36d236fac0828"
	password := "asaodUU0CLffNucp14KV3ymIdWhL97DPyXqevvBi"
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Shit is fucked up: %s", err)
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	      fmt.Println(err.Error())
	      panic(err)
	}

	 body := string(bytes)

	// Respond with user data
	response := Response{StatusCode: 200, Body: body}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
