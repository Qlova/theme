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
		fmt.Println("external widgets are not supported yet, sorry :/")
		return
	}

	os.Mkdir("theme", os.ModePerm)
	os.Mkdir("theme/"+widget, os.ModePerm)

	var path = "theme/" + widget + "/" + widget + ".go"
	if _, err := os.Stat(path); err != nil {
		var file, _ = os.Create(path)

		file.Write([]byte(`package ` + widget + `

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/` + widget + `"

type Widget struct {
	` + widget + `.Widget
}

func New() Widget {
	widget := ` + widget + `.New()

	//Add custom styles here.

	return  Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
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
