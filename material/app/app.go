package app

import (
	"github.com/qlova/seed"
	"github.com/qlova/theme/material"
)

type Options struct {
	Name       string
	ColorTheme material.ColorTheme
}

func New(options Options) *seed.App {
	var App = seed.NewApp(options.Name)

	if options.ColorTheme != (material.ColorTheme{}) {
		material.Color = options.ColorTheme
	}

	return App
}
