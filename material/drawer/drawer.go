package drawer

import (
	"github.com/qlova/seed"
)

type Seed struct {
	seed.Seed
}

func New(options Options) Seed {
	var Drawer = seed.New()

	Drawer.Require("material.css")
	Drawer.Require("material.js")

	Drawer.SetClass("mdl-layout__drawer")

	if options.Title != "" {
		var Title = seed.AddTo(Drawer)
		Title.SetTag("span")
		Title.SetText(options.Title)
		Title.SetClass("mdl-layout-title")
	}

	if len(options.Items) != 0 {
		var Naviagtion = seed.AddTo(Drawer)
		Naviagtion.SetTag("nav")
		Naviagtion.SetClass("mdl-navigation")
		for i := range options.Items {
			item := options.Items[i]
			var Text = seed.AddTo(Naviagtion)
			Text.SetText(item.Name)
			Text.SetTag("a")
			Text.SetClass("mdl-navigation__link")
			Text.OnClick(func(q seed.Script) {
				item.OnClick(q)
				q.Javascript(Drawer.Script(q).Element() + `.parentElement.MaterialLayout.toggleDrawer();`)
			})
		}
	}

	return Seed{Drawer}
}

func AddTo(parent seed.Interface, options Options) Seed {
	var Drawer = New(options)
	parent.Root().Add(Drawer)
	return Drawer
}
