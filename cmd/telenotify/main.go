package main

import (
	"context"
	"log"

	stdHttp "net/http"

	"github.com/joho/godotenv"

	"github.com/pokatomnik/telenotify/internal/util/http"

	// adapters
	aTelegramPackage "github.com/pokatomnik/telenotify/internal/adapters/telegram"

	// repositories
	rSkillsRepoPackage "github.com/pokatomnik/telenotify/internal/repositories/skills"

	// Cobra controllers
	cmdMCPRootControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp"
	cmdMCPHTTPControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp_http"
	cmdMCPStdIOControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/mcp_stdio"
	cmdNotifyControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/notify"
	cmdSkillRemovePackage "github.com/pokatomnik/telenotify/internal/controllers/cli/remove_skill"
	cmdRootControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/root"
	cmdSkillExistsPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/skill"
	cmdSkillInstallPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/write_skill"

	// MCP controllers
	mcpHTTPControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/mcp/http"
	mcpStdIOControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/mcp/stdio"

	// use cases
	ucNotifyUseCasePackage "github.com/pokatomnik/telenotify/internal/use_cases/notify"
	ucSkillExistsPackage "github.com/pokatomnik/telenotify/internal/use_cases/skill_exists"
	ucSkillRemovePackage "github.com/pokatomnik/telenotify/internal/use_cases/skill_remove"
	ucSkillInstallPackage "github.com/pokatomnik/telenotify/internal/use_cases/skill_write"
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

	skillsRepo := rSkillsRepoPackage.New()

	// use cases
	ucNotify := ucNotifyUseCasePackage.New(aTelegram)
	ucSkillExists := ucSkillExistsPackage.New(skillsRepo)
	ucSkillInstall := ucSkillInstallPackage.New(skillsRepo)
	ucSkillRemove := ucSkillRemovePackage.New(skillsRepo)

	// MCP controllers
	mcpStdIOController := mcpStdIOControllerPackage.New(ucNotify)
	mcpHTTPController := mcpHTTPControllerPackage.New(ucNotify, config)

	// cobra controllers
	cmdNotify := cmdNotifyControllerPackage.Notify(ucNotify)

	cmdMCPStdIO := cmdMCPStdIOControllerPackage.MCPStdIOController(mcpStdIOController)
	cmdMCPHTTP := cmdMCPHTTPControllerPackage.MCPHTTPController(mcpHTTPController)
	cmdMCPRoot := cmdMCPRootControllerPackage.NewMCPRootController(cmdMCPStdIO, cmdMCPHTTP)

	cmdSkillInstall := cmdSkillInstallPackage.NewWriteSkillController(ucSkillInstall)
	cmdSkillRemove := cmdSkillRemovePackage.NewRemoveSkillController(ucSkillRemove)
	cmdSKillExists := cmdSkillExistsPackage.SkillController(ucSkillExists, cmdSkillInstall, cmdSkillRemove)

	cmdRoot := cmdRootControllerPackage.RootController(cmdNotify, cmdMCPRoot, cmdSKillExists)

	err = cmdRoot.ExecuteContext(appContext)

	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}
}
