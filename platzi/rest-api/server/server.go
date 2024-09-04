package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Printf("Server running on port %s\n", b.Config().Port)
	if err := http.ListenAndServe(":"+b.Config().Port, b.router); err != nil {
		log.Fatal("Server error: ", err)
	}
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DataBaseURL == "" {
		return nil, errors.New("database url is required")
	}
	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}
