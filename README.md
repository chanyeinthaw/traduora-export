# Traduora Exporter

Export translations from traduora.

# Requirements

- traduora host
- client credentials

> client credentials can be retrieved from traduora > project > API Keys

# Building from source

```
go mod tidy
go build .
```

# Using a pre-build binary

Download the latest binary from [releases](https://github.com/chanyeinthaw/traduora-export/releases).

# Usage

Just run the command `traduora-export`. If a configuration file is not already set up, it will generate one. The tool will then download the translation files and save them under the configured output directory.
