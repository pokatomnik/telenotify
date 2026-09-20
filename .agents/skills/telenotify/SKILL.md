---
name: telenotify
description: Send Telegram notifications with the Telenotify CLI. Use when an agent needs to compose and execute a notification command, choose a request timeout, or troubleshoot CLI usage.
---

# Telenotify CLI

Use Telenotify to send a single message to the user.

## Command

The notification command requires exactly one message argument:

```sh
telenotify notify "Deployment completed successfully"
```

Preserve the user's message text exactly unless the user explicitly asks for reformatting. Quote the message so whitespace, punctuation, and shell metacharacters are passed as one argument.

Send exactly one notification per request unless the user explicitly asks for multiple messages.

## Timeout

Each request has a 30-second timeout by default. Override it with `--timeout` or `-t`, using a Go duration:

```sh
telenotify notify --timeout 10s "This request has a shorter timeout"
telenotify notify -t 2m "Allow more time for this request"
```

A timeout limits how long the command waits for the request. It does not retry a failed notification. Do not claim delivery unless the command exits successfully.

## Help

Use the built-in help when command syntax or options are uncertain:

```sh
telenotify --help
telenotify notify --help
```

Running `telenotify` without a subcommand displays the root help.

## Examples

```sh
telenotify notify "Build #42 passed"
telenotify notify -t 60s "The backup job finished"
```
