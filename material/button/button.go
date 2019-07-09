package button

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/button"

	_ "github.com/qlova/theme/material"
)

type Seed struct {
	button.Seed
}

func New(options Options) Seed {
	var Button = button.New()

	Button.Require("material.css")
	Button.Require("material.js")

	Button.SetClass("mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect")
	Button.SetText(options.Text)

	return Seed{Button}
}

func AddTo(parent seed.Interface, options Options) Seed {
	var Button = New(options)
	parent.Root().Add(Button)
	return Button
}
