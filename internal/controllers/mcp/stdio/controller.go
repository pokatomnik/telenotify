package http

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	tools "github.com/pokatomnik/telenotify/internal/controllers/mcp/tools/notify"
	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
	"github.com/pokatomnik/telenotify/internal/prompts"
)

type notifier interface {
	Notify(ctx context.Context, req notifreq.NotificationRequest) error
}

type MCPHTTPRunner struct {
	notifier notifier
}

func New(notifier notifier) MCPHTTPRunner {
	return MCPHTTPRunner{
		notifier: notifier,
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

func (m MCPHTTPRunner) Run(ctx context.Context) error {
	server, err := m.getServer()
	if err != nil {
		return err
	}

	return server.Run(ctx, nil)
}
