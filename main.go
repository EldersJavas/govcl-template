package main

import (
	"./form"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.SetTitle("project")
	vcl.Application.SetScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form.Mainform)
	vcl.Application.Run()
}
