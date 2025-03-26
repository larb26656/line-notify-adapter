package line_bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/larb26656/line-notify-adapter/errs"
)

type LineBotService interface {
	SendMessage(accessToken string, to string, messageText string) error
}

type lineBotService struct {
}

func NewLineBotService() LineBotService {
	return &lineBotService{}
}

func (s *lineBotService) SendMessage(accessToken string, to string, messageText string) error {
	url := "https://api.line.me/v2/bot/message/push"

	requestBody := RequestBody{
		To: to,
		Messages: []Message{
			{
				Type: "text",
				Text: messageText,
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.Join(errs.ErrInvalidAccessToken, fmt.Errorf("failed to send message, status: %d", resp.StatusCode))
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message, status: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
