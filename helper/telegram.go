package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendMessageBotTelegram(token, chatId, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	httpClient := new(http.Client)

	body := strings.NewReader(fmt.Sprintf(`
	{
		"chat_id":"%s",
		"text": "%s"
	}
	`, chatId, message))

	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")

	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	resBytes, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	if httpRes.StatusCode >= 400 {
		return fmt.Errorf("status code: %d. response: %s", httpRes.StatusCode, string(resBytes))
	}

	return nil
}
