package layout

import (
	"github.com/qlova/seed"

	_ "github.com/qlova/theme/material"
)

type Seed struct {
	seed.Seed
}

func New(options Options) Seed {
	var Layout = seed.New()

	Layout.Require("material.css")
	Layout.Require("material.js")

	Layout.SetClass("mdl-layout mdl-js-layout mdl-layout--fixed-header")

	Layout.Add(options.AppBar)

	return Seed{Layout}
}

func AddTo(parent seed.Interface, options Options) Seed {
	var Layout = New(options)
	parent.Root().Add(Layout)
	return Layout
}
