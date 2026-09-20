package main

import (
	"context"
	"log"

	stdHttp "net/http"

	"github.com/joho/godotenv"

	"github.com/pokatomnik/telenotify/internal/util/http"

	// adapters
	aTelegramPackage "github.com/pokatomnik/telenotify/internal/adapters/telegram"

	// Cobra controllers
	cmdMCPRootControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp"
	cmdMCPHTTPControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp_http"
	cmdMCPStdIOControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp_stdio"
	cmdNotifyControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/notify"
	cmdRootControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/root"

	// MCP controllers
	mcpHTTPControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/mcp/http"
	mcpStdIOControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/mcp/stdio"

	// use cases
	ucNotifyUseCasePackage "github.com/pokatomnik/telenotify/internal/use_cases/notify"
)

func main() {
	// Skip godotenv error
	godotenv.Load(".env")

	config, err := ParseConfig()
	if err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	proxyURL, hasProxy, err := config.ProxyURL()
	if err != nil {
		log.Fatal("Failed to read proxy URL: ", err)
	}

	client := stdHttp.DefaultClient
	if hasProxy {
		client = http.NewClient(proxyURL)
	}

	// main context
	appContext := context.Background()

	// adapters
	aTelegram := aTelegramPackage.New(config, *client)

	// use cases
	ucNotify := ucNotifyUseCasePackage.New(aTelegram)

	// MCP controllers
	mcpStdIOController := mcpStdIOControllerPackage.New(ucNotify)
	mcpHTTPController := mcpHTTPControllerPackage.New(ucNotify, config)

	// cobra controllers
	cmdNotify := cmdNotifyControllerPackage.Notify(ucNotify)
	cmdMCPStdIO := cmdMCPStdIOControllerPackage.MCPStdIOController(mcpStdIOController)
	cmdMCPHTTP := cmdMCPHTTPControllerPackage.MCPHTTPController(mcpHTTPController)
	cmdMCPRoot := cmdMCPRootControllerPackage.NewMCPRootController(cmdMCPStdIO, cmdMCPHTTP)
	cmdRoot := cmdRootControllerPackage.RootController(cmdNotify, cmdMCPRoot)

	err = cmdRoot.ExecuteContext(appContext)

	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}
}
