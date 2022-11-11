package http

import (
	"errors"
	"github.com/halilylm/url-shortner/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type URLHandler struct {
	urlUsecase domain.URLUsecase
}

func NewURLHandler(e *echo.Echo, urlUsecase domain.URLUsecase) {
	handler := URLHandler{urlUsecase}
	e.POST("/url/insert", handler.Generate)
	e.GET("/:id", handler.Redirect)
}

func (u *URLHandler) Generate(c echo.Context) error {
	var url domain.Url
	if err := c.Bind(&url); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"error":   true,
			"message": err.Error(),
		})
	}
	createdURL, err := u.urlUsecase.Generate(&url)
	if err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.Code, httpErr)
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusInternalServerError,
			"error":   true,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, createdURL)
}

func (u *URLHandler) Redirect(c echo.Context) error {
	id := c.Param("id")
	url, err := u.urlUsecase.Redirect(id)
	if err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.Code, httpErr)
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := c.Redirect(http.StatusFound, url.LongURL); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": url})
}
