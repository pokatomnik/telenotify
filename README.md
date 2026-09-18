# Telenotify

Telenotify is a command-line application for sending messages to a Telegram chat through a Telegram bot.

## Requirements

- Go 1.27 or later
- A Telegram bot token
- The target Telegram chat ID

## Configuration

Set the following environment variables before running the application:

- `TELENOTIFY_BOT_TOKEN` — Telegram bot token.
- `TELENOTIFY_CHAT_ID` — ID of the chat that should receive messages.
- `FUCK_RKN_PROXY` — optional proxy URL for Telegram API requests.

You can also place these variables in a `.env` file in the working directory.

## Usage

Send a message with:

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

## Development

Run linting and tests:

```sh
make lint
make test
```

Build for the current platform:

```sh
make build
```

Build binaries for all supported platforms:

```sh
make build-all
```

Install the application using Go:

```sh
make install
```

Remove generated binaries:

```sh
make clean
```

The build process uses path trimming and linker flags to reduce binary size.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
