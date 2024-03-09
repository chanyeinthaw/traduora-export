package main

import (
	"github.com/chanyeinthaw/traduora-sync/auth"
	"github.com/chanyeinthaw/traduora-sync/config"
	"github.com/chanyeinthaw/traduora-sync/traduora"
)

func main() {
	config.Read()
	auth.Init()
	traduora.ValidateLocales()

	traduora.ExportTranslations()
}
