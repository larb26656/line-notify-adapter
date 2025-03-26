package notify

import (
	"errors"
	"fmt"
	"strings"

	"github.com/larb26656/line-notify-adapter/errs"
	"github.com/larb26656/line-notify-adapter/external/line_bot"
)

type NotifyService interface {
	SendNotify(token string, message string) (*SendNotifyRes, error)
}

type notifyService struct {
	LineBotService line_bot.LineBotService
}

func NewNotifyService(lineBotService line_bot.LineBotService) NotifyService {
	return &notifyService{
		LineBotService: lineBotService,
	}
}

func (s *notifyService) extractAuthorization(token string) (string, string, error) {
	parts := strings.Split(token, "_targetKey_")

	if len(parts) != 2 {
		return "", "", errors.Join(errs.ErrInvalidAccessToken, fmt.Errorf("invalid token pattern"))
	}

	channelAccessToken := parts[0]
	targetToken := parts[1]

	return channelAccessToken, targetToken, nil
}

func (s *notifyService) SendNotify(token string, message string) (*SendNotifyRes, error) {
	channelAccessToken, targetToken, err := s.extractAuthorization(token)

	if err != nil {
		return nil, err
	}

	err = s.LineBotService.SendMessage(channelAccessToken, targetToken, message)

	if err != nil {
		return nil, err
	}

	return &SendNotifyRes{
		Status:  200,
		Message: "ok",
	}, nil
}
