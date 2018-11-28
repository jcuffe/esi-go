package main

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type SystemCostIndices []System

type System struct {
	CostIndices	[]CostIndex	`json:"cost_indices"`
	SolarSystemID	int		`json:"solar_system_id"`
}

type CostIndex struct {
	Activity	string	`json:"activity"`
	CostIndex	float64	`json:"cost_index"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	query := request.QueryStringParameters
	systemID, err := strconv.Atoi(query["systemID"])
	if err != nil {
		systemID = 30004759
	}

	fmt.Println("System ID:", systemID)

	resp, err := http.Get("https://esi.evetech.net/latest/industry/systems")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var systems SystemCostIndices
	if err := json.Unmarshal(bytes, &systems); err != nil {
		panic(err)
	}

	var buildSystem System
	for _, system := range systems {
		if system.SolarSystemID == systemID {
			buildSystem = system
		}
	}

	bytes, err = json.Marshal(buildSystem)
	if err != nil {
		panic(err)
	}

	body := string(bytes)
	response := Response{
		Body: body,
		StatusCode: 200,
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
