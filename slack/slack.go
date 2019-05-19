package slack

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Post message to slack
func Post(message string) error {
	URL := "https://slack.com/api/chat.postMessage"
	values := url.Values{}
	values.Add("channel", os.Getenv("SLACK_CHANNEL"))
	values.Add("text", message)

	req, err := http.NewRequest("POST", URL, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SLACK_TOKEN"))

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
