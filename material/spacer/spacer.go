//Provides an spacer that will expand to use up empty space.
package spacer

import (
	"github.com/qlova/seed"

	_ "github.com/qlova/theme/material"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Spacer = seed.New()

	Spacer.SetClass("mdl-layout-spacer")

	return Seed{Spacer}
}

func AddTo(parent seed.Interface) Seed {
	var Spacer = New()
	parent.Root().Add(Spacer)
	return Spacer
}
