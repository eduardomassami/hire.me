package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardomassami/hime.me/config"
	"github.com/eduardomassami/hime.me/domain"
	"github.com/eduardomassami/hime.me/handlers"
	"github.com/eduardomassami/hime.me/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	conf, e := config.NewConfig("./config/config.yaml")
	if e != nil {
		log.Panic(e)
	}

	repo, err := repository.NewURLRepositoryDb(conf.Database.User, conf.Database.Password, conf.Database.Address, conf.Database.Port, conf.Database.Name, conf.Database.Timeout)
	if err != nil {
		log.Panic(err)
	}

	service := domain.NewURLService(repo)
	handler := handlers.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/retrieve/{alias}", handler.Get)
	r.Get("/retrieve", handler.GetMostUsed)
	r.Put("/create/alias", handler.SaveWithCustomAlias)
	r.Put("/create", handler.SaveNoCustomAlias)
	// r.Put("/create?url={url}&CUSTOM_ALIAS={alias}", handler.SaveWithCustomAlias)
	// r.Put("/create?url={url}", handler.SaveNoCustomAlias)

	log.Print(fmt.Sprintf("Starting server on %s:%s ...", conf.Server.Host, conf.Server.Port))
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
