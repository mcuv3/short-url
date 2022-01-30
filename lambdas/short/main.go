package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mcuv3/demo/internal/dynamo"
)

type Payload struct {
	PathParameters map[string]string `json:"pathParameters"`
}

func HandleRequest(ctx context.Context, payload Payload) (events.APIGatewayProxyResponse, error) {

	repo, err := dynamo.NewRepository()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	v, ok := payload.PathParameters["short"]

	if !ok || v == "" {
		return events.APIGatewayProxyResponse{}, errors.New("short parameter is required")
	}

	link, err := repo.GetCustomerByShort(v)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 301,
		Headers: map[string]string{
			"Location": link.FullURL,
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
