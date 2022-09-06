package config

import (
	"io/fs"
)

type Assets struct {
	Logo            string
	EmbedMigrations fs.FS
	TemplateSQL     fs.FS
}

func NewAssets(logo string, embedMigrations fs.FS, templateSQL fs.FS) *Assets {
	return &Assets{Logo: logo, EmbedMigrations: embedMigrations, TemplateSQL: templateSQL}
}
