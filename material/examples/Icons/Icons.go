package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/column"
	"github.com/qlova/seeds/expander"
	"github.com/qlova/seeds/row"
	"github.com/qlova/seeds/spacer"
	"github.com/qlova/seeds/text"

	"github.com/qlova/theme/material/app"
	"github.com/qlova/theme/material/appbar"
	"github.com/qlova/theme/material/layout"

	"github.com/qlova/theme/material/icon"
)

func main() {
	var App = app.New(app.Options{
		Name: "Icons",
	})

	var Layout = layout.AddTo(App, layout.Options{
		AppBar: appbar.New(appbar.Options{
			Title: "Icons",
		}),
	})

	expander.AddTo(Layout)

	spacer.AddTo(Layout, seed.Em*2)

	var Row = row.AddTo(Layout)
	Row.Wrap()
	Row.Center()
	Row.AlignChildren(0)

	for _, ico := range icon.All {
		var Column = column.AddTo(Row)
		Column.SetSize(10*seed.Em, 8*seed.Em)
		Column.Clip()

		var s = ico
		Column.OnClick(func(q seed.Script) {

			q.SetClipboard(q.String(s))
		})

		icon.AddTo(Column, ico).SetSize(4 * seed.Em)
		text.AddTo(Column, ico)
	}
	expander.AddTo(Layout)

	App.Launch()
}
