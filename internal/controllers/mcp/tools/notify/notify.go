package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
)

type notifier interface {
	Notify(ctx context.Context, req notifreq.NotificationRequest) error
}

type input struct {
	Message string `json:"message" jsonschema:"message to send"`
}

type output struct {
	Ok    bool   `json:"ok" jsonschema:"Ok"`
	Error string `json:"error" jsonschema:"error message"`
}

func ToolNotify(notifier notifier) (mcp.Tool, mcp.ToolHandlerFor[input, output], error) {
	tool := mcp.Tool{
		Name:        "notifier",
		Description: "Send message to user",
	}

	var handler mcp.ToolHandlerFor[input, output] = func(
		ctx context.Context,
		req *mcp.CallToolRequest,
		input input,
	) (
		*mcp.CallToolResult,
		output,
		error,
	) {
		err := notifier.Notify(ctx, notifreq.NotificationRequest{
			NotificationText: input.Message,
		})
		if err != nil {
			return &mcp.CallToolResult{IsError: true}, output{Error: err.Error()}, nil
		}
		return nil, output{Ok: true}, nil
	}

	return tool, handler, nil
}
