package stdio

import "context"

type MCPStdioRunner struct{}

func New() MCPStdioRunner {
	return MCPStdioRunner{}
}

func (r MCPStdioRunner) Run(ctx context.Context) error {
	panic("Not implemented")
}
