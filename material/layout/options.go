package layout

import (
	"github.com/qlova/theme/material/appbar"
	"github.com/qlova/theme/material/content"
	"github.com/qlova/theme/material/drawer"
)

type Options struct {
	AppBar  appbar.Seed
	Drawer  drawer.Seed
	Content content.Seed
}
