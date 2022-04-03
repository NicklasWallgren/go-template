package config

import (
	"embed"
	"io/fs"
)

type Assets struct {
	Logo            string
	EmbedMigrations fs.FS
}

func NewAssets(logo string, embedMigrations embed.FS) *Assets {
	return &Assets{Logo: logo, EmbedMigrations: embedMigrations}
}
