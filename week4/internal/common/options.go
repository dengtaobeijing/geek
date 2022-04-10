package common

import (
	"context"
	"log"
	"net/http"
)

type options struct {
	ctx     context.Context
	logger  *log.Logger
	servers []*http.Server
	name    string
}

type Option func(*options)

func Ctx(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func Logger(logger *log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func Servers(srv ...*http.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}

func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}
