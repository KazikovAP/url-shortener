package in_memory

import (
	"time"

	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

type URLRepository struct {
	db              map[string]model.URLInfo
	constraintCheck map[string]struct{}
	bigsirealValue  int64
}

func NewRepository() *URLRepository {
	return &URLRepository{
		db:              make(map[string]model.URLInfo),
		constraintCheck: make(map[string]struct{}),
		bigsirealValue:  0,
	}
}

func (r *URLRepository) InsertURLInfo(urlInfoModel model.URLInfo) (int64, error) {
	if _, ok := r.constraintCheck[urlInfoModel.URL]; ok {
		return 0, repository.ErrURLExists
	} else {
		r.constraintCheck[urlInfoModel.URL] = struct{}{}
	}

	r.bigsirealValue++

	urlInfoModel.CreatedAt = time.Now()
	urlInfoModel.ID = r.bigsirealValue

	r.db[urlInfoModel.Alias] = urlInfoModel

	return urlInfoModel.ID, nil
}

func (r *URLRepository) GetURLInfoByAlias(alias string) (model.URLInfo, error) {
	urlInfo, ok := r.db[alias]
	if !ok {
		return model.URLInfo{}, repository.ErrURLNotFound
	}

	return urlInfo, nil
}
