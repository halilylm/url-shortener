package server

import (
	"github.com/couchbase/gocb/v2"
	"github.com/halilylm/url-shortner/url/delivery/http"
	"github.com/halilylm/url-shortner/url/repository/couchbase"
	"github.com/halilylm/url-shortner/url/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Server) setupRoutes() {
	cluster, err := gocb.Connect("couchbase://db", gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	})
	if err != nil {
		logrus.Fatal(err)
	}
	repo := couchbase.NewUrlRepository(cluster)
	ucase := usecase.NewURLUsecase(repo, s.cfg)
	http.NewURLHandler(s.e, ucase)
}
