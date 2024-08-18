package main

import (
	"fmt"
	"os"
	"svgTD"
)

func main() {
	var width, height = 500, 500
	svg := svg.NewWH(width, height)
	svg.Circle().Attrs(width/2, height/2, 100).
		Text("Hello World").Attrs(width/2, height/2).
		Style("text-anchor", "middle").
		Style("font-size", "30px").
		Style("fill", "white")
	svg.Pretty(true)
	if b, err := svg.Marshal(); err != nil {
		fmt.Println("cannot generate SVG %w", err)
		os.Exit(1)
	} else {
		fmt.Println(string(b))
	}
}
