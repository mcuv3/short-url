package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mcuv3/demo/internal/dynamo"
	"github.com/mcuv3/demo/internal/link"
	"github.com/mcuv3/demo/internal/models"
	"github.com/mcuv3/demo/internal/validation"
)

type Payload struct {
	URL string `json:"u"`
}

type Response struct {
	Alias string `json:"a"`
}

func HandleRequest(ctx context.Context, payload Payload) (Response, error) {

	if ok := validation.ValidateURL(payload.URL); !ok {
		return Response{}, errors.New("invalid url")
	}

	urlShort := link.FromURL(payload.URL)
	alias := link.Encode(urlShort)

	dyn, err := dynamo.NewRepository()
	if err != nil {
		return Response{}, err
	}

	err = dyn.SaveLink(&models.Link{
		FullURL: payload.URL,
		Short:   alias,
		Title:   "to_implement",
	})

	if err != nil {
		fmt.Printf("error creating link: %v", err)
	}

	return Response{
		Alias: fmt.Sprintf("https://url.mcuve.com/%s", alias),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
