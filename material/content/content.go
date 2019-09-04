package content

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Content = seed.New()
	Content.SetClass(`mdl-layout__content`)
	return Seed{Content}
}

func AddTo(parent seed.Interface) Seed {
	var Content = New()
	parent.Root().Add(Content)
	return Content
}
