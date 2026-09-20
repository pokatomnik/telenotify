package main

import (
	"net/url"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	EnvChatID      string `env:"TELENOTIFY_CHAT_ID,required,notEmpty"`
	EnvBotToken    string `env:"TELENOTIFY_BOT_TOKEN,required,notEmpty"`
	EnvProxyURL    string `env:"FUCK_RKN_PROXY"`
	EnvMCPHTTPHost string `env:"MCP_HTTP_HOST"`
}

func ParseConfig() (Config, error) {
	return env.ParseAs[Config]()
}

func (c Config) Token() string {
	return c.EnvBotToken
}

func (c Config) ChatID() string {
	return c.EnvChatID
}

func (c Config) HTTPAddress() string {
	return c.EnvMCPHTTPHost
}

func (c Config) ProxyURL() (url *url.URL, hasProxy bool, err error) {
	if c.EnvProxyURL == "" {
		return nil, false, nil
	}

	urlParsed, err := url.Parse(c.EnvProxyURL)
	if err != nil {
		return nil, false, err
	}

	return urlParsed, true, nil
}
