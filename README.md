# scriptPixelPerfect

scriptPixelPerfect add data set in modelsList in database

## Prerequisites

- GoLang SDK (^1.20.2)
- MSSQL Server

## Config

### File

ServerRest supports configuration file in `JSON` format.

The location for the configuration file is at the system root under .config.yaml file.

Please see below and/or in the .config.example.yaml for examples of valid configuration files

**JSON :**

```json
{
    "server": "serverName",
    "database": "databaseName",
    "user": "User",
    "password": "Password",
    "port": 3306,
    "timeout": "30"
}
```

## Developing

1. Download project with `git clone git@github.com:EllVodka/pixel-perfect-script.git` or `git clone https://github.com/EllVodka/pixel-perfect-script.git`
2. Run project with `make run`

## Building

You can run the commande `make build`.
You can build for specific os (windows,linux,darwin) with commande  `make specific-os`.
