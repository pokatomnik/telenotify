package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type listener interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type HTTPServerExt struct {
	server listener
}

func NewHTTPServerExt(orig *http.Server) *HTTPServerExt {
	return &HTTPServerExt{server: orig}
}

func (s *HTTPServerExt) ListenAndServeContext(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	lch := make(chan error, 1)
	sch := make(chan error, 1)

	eg.Go(func() error {
		err := s.server.ListenAndServe()
		lch <- err
		return err
	})

	eg.Go(func() error {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Second*5,
		)
		defer cancel()

		err := s.server.Shutdown(stopCtx)
		sch <- err
		return err
	})

	eg.Wait()

	return errors.Join(<-lch, <-sch)
}
