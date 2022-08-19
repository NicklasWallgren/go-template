package config

import (
	"io/fs"
)

type Assets struct {
	Logo            string
	EmbedMigrations fs.FS
}

func NewAssets(logo string, embedMigrations fs.FS) *Assets {
	return &Assets{Logo: logo, EmbedMigrations: embedMigrations}
}
