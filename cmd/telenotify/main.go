package main

import (
	"context"
	"log"

	stdHttp "net/http"

	"github.com/joho/godotenv"
	"github.com/pokatomnik/telenotify/internal/adapters/telegram"
	ucNotifyControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/notify"
	ucRootControllerPackage "github.com/pokatomnik/telenotify/internal/controllers/cli/root"
	ucNotifyUseCasePackage "github.com/pokatomnik/telenotify/internal/use_cases/notify"
	"github.com/pokatomnik/telenotify/internal/util/http"
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
	aTelegram := telegram.New(config, *client)

	// use cases
	ucNotify := ucNotifyUseCasePackage.New(aTelegram)

	// controllers
	cmdNotify := ucNotifyControllerPackage.Notify(ucNotify)
	cmdRoot := ucRootControllerPackage.RootController(cmdNotify)

	err = cmdRoot.ExecuteContext(appContext)

	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}
}
