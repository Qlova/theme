package material

import "image/color"
import "github.com/qlova/seed"

func init() {
	seed.Embed("/material.js", []byte(JS))
	seed.Embed("/material.css", []byte(CSS))
}

type ColorTheme struct {
	Primary, PrimaryVariant,
	Secondary, SecondaryVariant,
	Background, Surface, Error,
	OnPrimary, OnSecondary,
	OnBackground, OnSurface, OnError color.Color
}

var Color ColorTheme = DefaultTheme

var DefaultTheme = ColorTheme{
	Primary:        seed.Hex("#6200EE"),
	PrimaryVariant: seed.Hex("#3700B3"),

	Secondary:        seed.Hex("#03DAC6"),
	SecondaryVariant: seed.Hex("#018786"),

	Background: seed.White,
	Surface:    seed.White,
	Error:      seed.Hex("#B00020"),

	OnPrimary:    seed.White,
	OnSecondary:  seed.Black,
	OnBackground: seed.Black,
	OnSurface:    seed.Black,
	OnError:      seed.White,
}
