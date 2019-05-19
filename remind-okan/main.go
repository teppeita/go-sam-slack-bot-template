package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"remind_okan/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	errNotFirst   = errors.New("It's not a first request")
	errNotMention = errors.New("It's not a mention to app")
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Headers["X-Slack-Retry-Num"] != "" {
		return events.APIGatewayProxyResponse{}, errNotFirst
	}

	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)
	slackReq := new(postFromSlack)
	if err := json.Unmarshal(jsonBytes, &slackReq); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if slackReq.Challenge != "" {
		return events.APIGatewayProxyResponse{
			Body:       slackReq.Challenge,
			StatusCode: 200,
		}, nil
	}
	if slackReq.Event.Type != "app_mention" {
		return events.APIGatewayProxyResponse{}, errNotMention
	}

	// TODO: メッセージの組み立て
	msg := os.Getenv("MSG")

	// NOTE: slack通知
	err := slack.Post(msg)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Success"),
		StatusCode: 200,
	}, nil
}

type postFromSlack struct {
	Event struct {
		Type string
	}
	Challenge string
}
