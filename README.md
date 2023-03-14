`baton-jumpcloud` is a connector for JumpCloud built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It communicates with the JumpCloud API to sync data about users, groups, roles, and apps.

Check out [Baton](https://github.com/conductorone/baton) to learn more the project in general.

# Getting Started

## Prerequisites
- API key from JumpCloud
- (Optional) The Org ID for the JumpCloud account

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-jumpcloud
baton-jumpcloud
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_API_KEY=JumpCloudAPIKey ghcr.io/conductorone/baton-jumpcloud:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-jumpcloud/cmd/baton-jumpcloud@main

baton resources
```

# Data Model

`baton-jumpcloud` will pull down information about the following JumpCloud resources:

- Users
- Admin Users
- Groups
- Roles
- Apps

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-jumpcloud` Command Line Usage

```
baton-jumpcloud

Usage:
  baton-jumpcloud [flags]
  baton-jumpcloud [command]

Available Commands:
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --api-key string      The API key for the JumpCloud account.  ($BATON_API_KEY)
  -f, --file string         The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                help for baton-jumpcloud
      --log-format string   The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string    The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
      --org-id string       The Org ID for the JumpCloud account (optional). ($BATON_ORG_ID)
  -v, --version             version for baton-jumpcloud

Use "baton-jumpcloud [command] --help" for more information about a command.
```
