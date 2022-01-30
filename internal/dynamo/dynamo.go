package dynamo

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/mcuv3/demo/internal/models"
)

type DynamoCustomerRepository struct {
	db *dynamodb.DynamoDB
}

func NewRepository() (*DynamoCustomerRepository, error) {
	scs, err := session.NewSession()
	scs.Config.Region = aws.String("us-east-1")
	if err != nil {
		return nil, err
	}

	dy := dynamodb.New(scs)
	return &DynamoCustomerRepository{
		db: dy,
	}, nil
}

func (d *DynamoCustomerRepository) GetCustomerByShort(short string) (*models.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := d.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("url"),
		Key: map[string]*dynamodb.AttributeValue{
			"short": {
				S: aws.String(short),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(*res.Item["id"].S)
	if err != nil {
		return nil, err
	}

	return &models.Link{
		ID:      id,
		Short:   *res.Item["short"].S,
		FullURL: *res.Item["full_url"].S,
	}, nil
}

func (d *DynamoCustomerRepository) SaveLink(customer *models.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := uuid.New()

	_, err := d.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("url"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id.String()),
			},
			"short": {
				S: aws.String(customer.Short),
			},
			"full_url": {
				S: aws.String(customer.FullURL),
			},
			"created_at": {
				S: aws.String(time.Now().Format(time.RFC3339)),
			},
		},
	})

	if err == nil {
		return err
	}

	return nil
}
