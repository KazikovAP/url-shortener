package url

import (
	"fmt"
	"math/rand"
	"time"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

const aliasLength = 10

type Service struct {
	urlInfoRepository repository.URLInfoRepository
}

func NewService(urlInfoRepository repository.URLInfoRepository) *Service {
	return &Service{
		urlInfoRepository: urlInfoRepository,
	}
}

func (s *Service) SaveURL(url string) (string, error) {
	alias := newRandomString(aliasLength)

	urlInfo := model.URLInfo{
		URL:   url,
		Alias: alias,
	}

	_, err := s.urlInfoRepository.InsertURLInfo(urlInfo)
	if err != nil {
		return "", fmt.Errorf("urlInfoRepository.InsertURLInfo : %w", err)
	}

	return alias, nil
}

func (s *Service) GetURL(alias string) (string, error) {
	urlInfo, err := s.urlInfoRepository.GetURLInfoByAlias(alias)
	if err != nil {
		return "", fmt.Errorf("urlInfoRepository.InsertURLInfo : %w", err)
	}

	return urlInfo.URL, nil
}

func newRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789_")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
