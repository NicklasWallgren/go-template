package config

import (
	"io/fs"
)

type AssetsConfig struct {
	Logo            string
	EmbedMigrations fs.FS
	TemplateSQL     fs.FS
}

func NewAssetsConfig(logo string, embedMigrations fs.FS, templateSQL fs.FS) *AssetsConfig {
	return &AssetsConfig{Logo: logo, EmbedMigrations: embedMigrations, TemplateSQL: templateSQL}
}
