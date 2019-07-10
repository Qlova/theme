//Provides icons from the material icons font.
package icon

import (
	"github.com/qlova/seed"

	_ "github.com/qlova/theme/material"
)

type Seed struct {
	seed.Seed
}

func New(icon string) Seed {
	var Icon = seed.New()

	Icon.SetTag("i")
	Icon.SetClass("material-icons")
	Icon.SetText(icon)

	return Seed{Icon}
}

func AddTo(parent seed.Interface, icon string) Seed {
	var Icon = New(icon)
	parent.Root().Add(Icon)
	return Icon
}

func (icon Seed) SetSize(size complex128) {
	icon.SetTextSize(size)
}
