package repository

import "url-shortener/internal/model"

type URLInfoRepository interface {
	InsertURLInfo(urlInfoModel model.URLInfo) (int64, error)
	GetURLInfoByAlias(alias string) (model.URLInfo, error)
}
