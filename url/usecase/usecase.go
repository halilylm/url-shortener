package usecase

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/halilylm/url-shortner/config"
	"github.com/halilylm/url-shortner/domain"
	"github.com/labstack/echo/v4"
	"github.com/speps/go-hashids/v2"
	"net/http"
	"time"
)

type urlUsecase struct {
	urlRepository domain.URLRepository
	cfg           *config.Config
}

func NewURLUsecase(urlRepo domain.URLRepository, cfg *config.Config) domain.URLUsecase {
	return &urlUsecase{
		urlRepository: urlRepo,
		cfg:           cfg,
	}
}

func (u *urlUsecase) Generate(url *domain.Url) (*domain.Url, error) {
	hd, err := hashids.NewWithData(hashids.NewData())
	if err != nil {
		return nil, &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}
	id, err := hd.Encode([]int{int(time.Now().Unix())})
	if err != nil {
		return nil, &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}
	url.ShortURL = fmt.Sprintf("http://%s:%d/%s", u.cfg.Host, u.cfg.Port, id)
	url.ID = id
	createdURL, err := u.urlRepository.Insert(url)
	if err != nil {
		return nil, &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}
	return createdURL, nil
}

func (u *urlUsecase) Redirect(id string) (*domain.Url, error) {
	url, err := u.urlRepository.GetByID(id)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, &echo.HTTPError{
				Code:     http.StatusNotFound,
				Message:  err.Error(),
				Internal: err,
			}
		}
		return nil, &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}
	return url, nil
}
