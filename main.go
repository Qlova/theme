/*

	Theme Qlovaseed applications!

	This is a tool to quickly create new themed widgets.

*/
package main

import (
	"fmt"
	"strings"

	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("theme [widget]")
		return
	}

	var widget = os.Args[1]

	if strings.Contains(widget, "/") {
		fmt.Println("external seeds are not supported yet, sorry :/")
		return
	}

	os.Mkdir("theme", os.ModePerm)
	os.Mkdir("theme/"+widget, os.ModePerm)

	var path = "theme/" + widget + "/" + widget + ".go"
	if _, err := os.Stat(path); err != nil {
		var file, _ = os.Create(path)

		var Widget = strings.Title(widget)

		file.Write([]byte(`package ` + widget + `

import "github.com/qlova/seed"
import "github.com/qlova/seeds/` + widget + `"

type Seed struct {
	` + widget + `.Seed
}

func New() Seed {
	var ` + Widget + ` = ` + widget + `.New()

	//Add custom styles here.

	return  Seed{` + Widget + `}
}

func AddTo(parent seed.Interface) Seed {
	var ` + Widget + ` = New()
	parent.Root().Add(` + Widget + `)
	return ` + Widget + `
}

			`))

		file.Close()

		fmt.Println("New themed widget created!")
		return
	} else {
		fmt.Println(err)
	}

	var editor = os.Getenv("EDITOR")
	if editor == "" {
		fmt.Println("Please configure $EDITOR env variable, pointing to your text editor.")
		return
	}

	exec.Command(editor, path).Run()
}
