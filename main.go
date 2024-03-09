package main

import (
	"github.com/chanyeinthaw/traduora-export/auth"
	"github.com/chanyeinthaw/traduora-export/config"
	"github.com/chanyeinthaw/traduora-export/traduora"
)

func main() {
	config.Read()
	auth.Init()
	traduora.ValidateLocales()

	traduora.ExportTranslations()
}
