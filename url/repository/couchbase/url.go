package couchbase

import (
	"github.com/couchbase/gocb/v2"
	"github.com/halilylm/url-shortner/domain"
)

type urlRepository struct {
	bucket *gocb.Cluster
}

func NewUrlRepository(buck *gocb.Cluster) domain.URLRepository {
	return &urlRepository{buck}
}

func (u *urlRepository) Insert(url *domain.Url) (*domain.Url, error) {
	insertQuery := "INSERT INTO `urls` (KEY, VALUE) VALUES ($1, {'long_url': $2, 'short_url': $3})"
	values := []any{url.ID, url.LongURL, url.ShortURL}
	_, err := u.bucket.Query(insertQuery, &gocb.QueryOptions{PositionalParameters: values})
	if err != nil {
		panic(err)
		return nil, err
	}
	return url, nil
}

func (u *urlRepository) GetByID(id string) (*domain.Url, error) {
	var foundURL domain.Url
	selectQuery := "SELECT u.* FROM urls.`_default`.`_default` u WHERE meta().id = $1;"
	rows, err := u.bucket.Query(selectQuery, &gocb.QueryOptions{PositionalParameters: []any{id}})
	if err != nil {
		return nil, err
	}
	if err := rows.One(&foundURL); err != nil {
		return nil, err
	}
	return &foundURL, nil
}
