package http

import (
	"cmp"
	"context"
	"errors"
	"net"
	stdHttp "net/http"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	tools "github.com/pokatomnik/telenotify/internal/controllers/mcp/tools/notify"
	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
	"github.com/pokatomnik/telenotify/internal/prompts"
	"golang.org/x/sync/errgroup"
)

const (
	defaultHTTPAddress = "127.0.0.1:8080"
	shutdownTimeout    = 5 * time.Second
)

type notifier interface {
	Notify(ctx context.Context, req notifreq.NotificationRequest) error
}

type config interface {
	HTTPAddress() string
}

type MCPHTTPRunner struct {
	notifier notifier
	config   config
}

func New(notifier notifier, config config) MCPHTTPRunner {
	return MCPHTTPRunner{
		notifier: notifier,
		config:   config,
	}
}

func (m MCPHTTPRunner) getServer() (*mcp.Server, error) {
	implementation := &mcp.Implementation{
		Name:        "telenotify",
		Title:       "Telenotify",
		Description: "Notification MCP server",
		WebsiteURL:  "https://github.com/pokatomnik/telenotify",
	}
	options := &mcp.ServerOptions{
		Instructions: prompts.MCPInstructions,
	}

	server := mcp.NewServer(implementation, options)

	toolNotify, handlerNotify, err := tools.ToolNotify(m.notifier)
	if err != nil {
		return nil, err
	}
	mcp.AddTool(server, &toolNotify, handlerNotify)

	return server, nil
}

func (m MCPHTTPRunner) getHandler() (stdHttp.Handler, error) {
	server, err := m.getServer()
	if err != nil {
		return nil, err
	}

	return mcp.NewStreamableHTTPHandler(func(*stdHttp.Request) *mcp.Server {
		return server
	}, nil), nil
}

func (m MCPHTTPRunner) getHTTPAddress() string {
	return cmp.Or(m.config.HTTPAddress(), string(defaultHTTPAddress))
}

func (m MCPHTTPRunner) Run(ctx context.Context) error {
	handler, err := m.getHandler()
	if err != nil {
		return err
	}

	httpServer := &stdHttp.Server{
		Addr:    m.getHTTPAddress(),
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	eg, serverCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return httpServer.ListenAndServe()
	})

	eg.Go(func() error {
		<-serverCtx.Done()
		cancelCtx, cancel := context.WithTimeout(
			context.Background(),
			shutdownTimeout,
		)
		defer cancel()

		return httpServer.Shutdown(cancelCtx)
	})

	err = eg.Wait()
	if errors.Is(err, stdHttp.ErrServerClosed) {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
