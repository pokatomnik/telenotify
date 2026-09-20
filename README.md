# Telenotify

Telenotify is a command-line application that sends messages to a Telegram chat through a Telegram bot. It can also expose the notification functionality as an MCP server over stdio or HTTP.

## Requirements

- Go 1.27 or later
- A Telegram bot token
- The target Telegram chat ID

## Configuration

Configuration is read from environment variables. The application also loads a `.env` file from the current working directory when it starts.

- `TELENOTIFY_BOT_TOKEN` — required Telegram bot token.
- `TELENOTIFY_CHAT_ID` — required ID of the chat that should receive messages.
- `FUCK_RKN_PROXY` — optional proxy URL for Telegram API requests.
- `MCP_HTTP_HOST` — optional address for the MCP HTTP server. The default is `127.0.0.1:8080`.

Example `.env` file:

```dotenv
TELENOTIFY_BOT_TOKEN=your-bot-token
TELENOTIFY_CHAT_ID=your-chat-id
# FUCK_RKN_PROXY=http://127.0.0.1:8080
# MCP_HTTP_HOST=127.0.0.1:8080
```

Do not commit `.env` or expose the bot token.

## Usage

Send a message with the default 30-second request timeout:

```sh
telenotify notify "Hello from Telenotify"
```

Set a custom request timeout with the `--timeout` or `-t` option:

```sh
telenotify notify --timeout 10s "This request has a shorter timeout"
```

Run the application without a command to display the built-in help:

```sh
telenotify
```

### MCP server

Start an MCP server over standard input/output:

```sh
telenotify mcp stdio
```

Start the streamable HTTP MCP server. It listens on `127.0.0.1:8080` by default, or on the address configured with `MCP_HTTP_HOST`:

```sh
telenotify mcp http
```

## Development

### Architecture

The application follows a layered architecture:

- `cmd/telenotify` — application entry point and environment configuration.
- `internal/controllers/cli` — Cobra commands for notifications and MCP server startup.
- `internal/controllers/mcp` — MCP transports over stdio and streamable HTTP.
- `internal/controllers/mcp/tools` — MCP tool definitions and request handling.
- `internal/use_cases` — application use cases, including message sending.
- `internal/adapters` — integrations with external services, including the Telegram API.
- `internal/entities` — domain entities and errors shared by the application layers.
- `internal/util` — shared technical utilities, such as the HTTP client.

The CLI and MCP controllers depend on use cases rather than calling Telegram directly. This keeps transport-specific code separate from notification logic and allows the same use case to be used by the CLI and both MCP transports.

Run linting and tests:

```sh
make lint
make test
```

Build for the current platform:

```sh
make build
```

Build binaries for all supported platforms (Linux, macOS, and Windows on amd64 and arm64):

```sh
make build-all
```

Install the application with `go install`:

```sh
make install
```

Remove generated binaries:

```sh
make clean
```

The build uses path trimming and linker flags to reduce binary size. Set `VERSION` to embed a version string, for example:

```sh
make build VERSION=1.0.0
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
