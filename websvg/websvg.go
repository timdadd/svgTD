// websvg draws SVG in a web server
//go:build !appengine

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	svg "svgTD"
)

// Based upon https://github.com/ajstarks/svgo/tree/master/websvg

const pageStart = `<!DOCTYPE html>
<html lang="en">
<body>
<h2>Web SVG</h2>
`
const pageEnd = `</body>
</html>`

var port = flag.String("port", "2003", "http service address")

func main() {
	flag.Parse()
	http.Handle("/circle/", http.HandlerFunc(circle))
	http.Handle("/rectangle/", http.HandlerFunc(rect))
	http.Handle("/arc/", http.HandlerFunc(arc))
	http.Handle("/text/", http.HandlerFunc(text))
	fmt.Printf("http://localhost:%s/circle/\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Println("ListenAndServe:", err)
	}
}

func present(w http.ResponseWriter, r *http.Request, s *svg.SVG, e *svg.Element) {
	setShapeStyle(e, r.URL.RawQuery)
	fmt.Println(r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/shape") {
		_ = s.Write(w)
		return
	}
	menu := menuSVG()
	//menu.G().SVG(s)
	menu.Pretty(true)
	s.Pretty(true)
	fmt.Println(menu.ToString())
	//w.Header().Set("Content-Type", "image/svg+xml")
	_, _ = fmt.Fprintf(w, "%s%s\n%s\n%s",
		pageStart, menu.ToString(), s.ToString(), pageEnd)
	//_ = menu.Write(w)
	//fmt.Fprintf(w, "<div>%s</div>", menu.ToString())
	//_ = s.Write(w)
}

// setShapeStyle is allowing users to put style items on the URL
// Example:http://localhost:2003/circle?fill=blue
func setShapeStyle(e *svg.Element, rawQuery string) {
	fmt.Println(e.Tag(), rawQuery)
	// Set the default style
	e.Style("fill", "rgb(127,0,0)")
	if m, err := url.ParseQuery(rawQuery); err == nil {
		for k, v := range m {
			e.Style(k, v[0])
		}
	}
	return
}

func circle(w http.ResponseWriter, r *http.Request) {
	s := svg.NewWH(500, 500) //, "Circle")
	e := s.Circle().Attrs(250, 250, 125).Element
	present(w, r, s, e)
}

func rect(w http.ResponseWriter, r *http.Request) {
	s := svg.NewWH(500, 500).Title("Rectangle")
	e := s.Rect().Attrs(250, 250, 100, 200).Element
	present(w, r, s, e)
}

// arc is a path
// test with http://localhost:2003/arc/?fill=cyan&stroke=red
func arc(w http.ResponseWriter, r *http.Request) {
	s := svg.NewWHT(500, 500, "Arc")
	p := s.Path().MoveTo(true).P(250, 250). // Move to absolute point 250,250
						Arc(true). // start arc using absolute positions
						Attrs(100, 100, 0, false, false, 100, 125).Element()
	present(w, r, s, p)
}

func text(w http.ResponseWriter, r *http.Request) {
	s := svg.NewWH(500, 500).Title("Text")
	t := s.TextXY(250, 250, "Hello World").
		Style("text-anchor", "middle").
		Style("font-size", "32px")
	present(w, r, s, t.Element)
}

func menuSVG() (s *svg.SVG) {
	var menuItems = []string{"Circle", "Rectangle", "Arc", "Text"}
	width, height := 500, 200
	var xOffset = 50
	var yOffset = 20

	var menuItemWidth = 80
	var verticalSpacing = menuItemWidth + 40 // + x where x is the width between the items
	var menuItemHeight = 40
	var LabelOffsetY = 25

	width = xOffset + len(menuItems)*verticalSpacing
	height = yOffset + menuItemHeight
	// Create the s canvas
	s = svg.NewWHT(width, height, "Menu")

	// Draw classes
	for i := range menuItems {
		var x = xOffset + i*verticalSpacing
		s.G().Transform(svg.Translate(x, yOffset)).
			Class("menuItem").
			Rect().
			Attrs(-menuItemWidth/2, 0, menuItemWidth, menuItemHeight).
			Style("fill", "#CCC")
	}

	// Draw class labels in the class boxes
	for i, mi := range menuItems {
		var x = xOffset + i*verticalSpacing
		s.G().
			Transform(svg.Translate(x, yOffset)).
			Class("labels").
			Append("a").Attr("href", "/%s", strings.ToLower(mi)).
			Text(mi).Style("text-anchor", "middle").
			AttrsD(0, LabelOffsetY)
	}
	return
}
