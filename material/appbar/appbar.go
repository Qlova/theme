package appbar

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/text"

	"github.com/qlova/theme/material"
	"github.com/qlova/theme/material/spacer"
)

type Seed struct {
	seed.Seed
}

func New(options Options) Seed {
	var AppBar = seed.New()

	AppBar.SetTag("header")

	AppBar.Require("material.css")
	AppBar.Require("material.js")

	AppBar.SetClass("mdl-layout__header")

	AppBar.SetColor(material.Color.Primary)

	var Row = seed.AddTo(AppBar)
	{
		Row.SetClass("mdl-layout__header-row")
		Row.SetColor(material.Color.Primary)

		//if no other items on appbar.
		Row.SetInnerSpacing(0, 0)
		spacer.AddTo(Row)

		if options.Title != "" {
			var Title = text.AddTo(Row, options.Title)
			Title.SetClass("mdl-layout-title")
			Title.SetColor(material.Color.OnPrimary)
		}

		spacer.AddTo(Row)
	}

	return Seed{AppBar}
}

func AddTo(parent seed.Interface, options Options) Seed {
	var AppBar = New(options)
	parent.Root().Add(AppBar)
	return AppBar
}
