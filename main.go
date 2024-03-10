package main

import (
	_ "embed"
	"github.com/chanyeinthaw/traduora-export/auth"
	"github.com/chanyeinthaw/traduora-export/config"
	"github.com/chanyeinthaw/traduora-export/traduora"
)

//go:embed traduora.example.yaml
var exampleConfig string

func main() {
	config.Read(exampleConfig)
	auth.Init()
	traduora.ValidateLocales()

	traduora.ExportTranslations()
}
