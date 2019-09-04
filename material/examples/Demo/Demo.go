package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/text"
	"github.com/qlova/theme/material/app"
	"github.com/qlova/theme/material/appbar"
	"github.com/qlova/theme/material/button"
	"github.com/qlova/theme/material/content"
	"github.com/qlova/theme/material/drawer"
	"github.com/qlova/theme/material/layout"
)

func main() {
	var App = app.New(app.Options{
		Name: "Demo",
	})

	var Content = content.New()
	var Home = Content.NewPage()
	var AnotherPage = Content.NewPage()

	Home.SetTransition(seed.Fade)
	{
		text.AddTo(Home, "This is a demo app, click the button to go to another page.")
		button.AddTo(Home, button.Options{
			Text: "Click me",
		}).OnClickGoto(AnotherPage)
	}

	AnotherPage.SetTransition(seed.Fade)
	{
		text.AddTo(AnotherPage, "This is another page")
	}

	layout.AddTo(App, layout.Options{
		AppBar: appbar.New(appbar.Options{
			Title: "Hello World",
		}),
		Drawer: drawer.New(drawer.Options{
			Title: "Demo App",
			Items: []drawer.Item{
				drawer.NewItem("Home", Home),
				drawer.NewItem("AnotherPage", AnotherPage),
			},
		}),
		Content: Content,
	})

	App.SetPage(Home)
	App.Launch()
}
