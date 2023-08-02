package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

var client = new(http.Client)

func Handle(ctx context.Context, event Event) {
	fmt.Printf("Received %v (message %s)", event, event.Message())
	msg := SlackMessage{Text: event.Message(), Username: Config.Username, Icon: Config.Icon}

	req, err := http.NewRequest(http.MethodPost, Config.Endpoint, msg.Serialized())
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[%s] %s", resp.Status, string(body))
}

func main() {
	lambda.Start(Handle)
}
