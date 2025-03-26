package notify

import (
	"errors"
	"strings"
)

type NotifyService interface {
	SendNotify(key string) (*SendNotifyRes, error)
}

type notifyService struct{}

func NewNotifyService() NotifyService {
	return &notifyService{}
}

func (s *notifyService) extractKey(key string) (string, string, error) {
	parts := strings.Split(key, "_targetKey_")

	if len(parts) != 2 {
		return "", "", errors.New("invalid key format")
	}

	channelAccessToken := parts[0]
	targetToken := parts[1]

	return channelAccessToken, targetToken, nil
}

func (s *notifyService) SendNotify(key string) (*SendNotifyRes, error) {
	// TODO implement line message client
	_, _, err := s.extractKey(key)

	if err != nil {
		return nil, err
	}

	return &SendNotifyRes{
		Status:  200,
		Message: "ok",
	}, nil
}
