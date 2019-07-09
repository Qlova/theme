package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/expander"

	"github.com/qlova/theme/material/app"
	"github.com/qlova/theme/material/appbar"
	"github.com/qlova/theme/material/button"
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
	var Button = button.AddTo(Layout, button.Options{
		Text: "Click Me!",
	})
	Button.SetWidth(12 * seed.Em)
	Button.Center()
	Button.OnClick(func(q seed.Script) {
		Button.Script(q).SetText(q.String("You clicked me"))
	})
	expander.AddTo(Layout)

	App.Launch()
}
