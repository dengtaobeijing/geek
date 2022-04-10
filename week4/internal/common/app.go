package common

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	ctx    context.Context
	cancel func()
	logger *log.Logger
	opts   options
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		logger: log.Default(),
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		logger: options.logger,
		opts:   options,
	}
}

func (a *App) Start() error {
	a.logger.Println("start app[" + a.opts.name + "]...")
	g, ctx := errgroup.WithContext(a.ctx)

	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done()
			return srv.Shutdown(a.ctx)
		})
		g.Go(func() error {
			return srv.ListenAndServe()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.cancel()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
