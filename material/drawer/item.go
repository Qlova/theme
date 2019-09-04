package drawer

import "github.com/qlova/seed"

//Item is a drawer item that can be clicked on.
type Item struct {
	Name    string
	OnClick func(q seed.Script)
}

//NewItem returns an item with the specified name and the page it goes to onclick.
func NewItem(name string, page seed.Page) Item {
	return Item{
		Name: name,
		OnClick: func(q seed.Script) {
			page.Script(q).Goto()
		},
	}
}
