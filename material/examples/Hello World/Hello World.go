package main

import (
	"github.com/qlova/seeds/expander"
	"github.com/qlova/seeds/text"

	"github.com/qlova/theme/material/app"
	"github.com/qlova/theme/material/appbar"
	"github.com/qlova/theme/material/layout"
)

func main() {
	var App = app.New(app.Options{
		Name: "Hello World",
	})

	var Layout = layout.AddTo(App, layout.Options{
		AppBar: appbar.New(appbar.Options{
			Title: "Hello World",
		}),
	})

	expander.AddTo(Layout)
	text.AddTo(Layout, "Hello World")
	expander.AddTo(Layout)

	App.Launch()
}
