package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"url-shortener/internal/model"
	"url-shortener/internal/repository"

	"github.com/jmoiron/sqlx"
)

type URLRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) GetURLInfoByAlias(alias string) (model.URLInfo, error) {
	res := urlInfo{}

	err := r.db.Get(
		&res,
		queryGetURLInfoByAlias,
		alias,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.URLInfo{}, repository.ErrURLNotFound
		}

		return model.URLInfo{}, fmt.Errorf("db.Get %w", err)
	}

	return res.ToModoel(), nil
}

func (r *URLRepository) InsertURLInfo(urlInfoModel model.URLInfo) (int64, error) {
	row := r.db.QueryRow(
		queryInsertURLInfo,
		urlInfoModel.URL,
		urlInfoModel.Alias,
	)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		// Проверяем на ошибку уникальности.
		if strings.Contains(err.Error(), "23505") {
			// Костыль, надо кастовать к pq.Error.
			// if pgErr, ok := err.(*pq.Error); ok {
			// if pgErr.Code == pq.ErrorCode("23505") {
			// }
			return 0, repository.ErrURLExists
		}

		return 0, fmt.Errorf("row.Scan: %w", err)
	}

	return id, nil
}
