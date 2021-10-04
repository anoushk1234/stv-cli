package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px:5,Py: 5, Type: "Classic", Color: uint(0xFF4500)})
	Box.Print("Happy Birthday", "\n(☞ﾟヮﾟ)☞ STEVE ☜(ﾟヮﾟ☜)\nfrom Anoushk aka Maverick")
}
