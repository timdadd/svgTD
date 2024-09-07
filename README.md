# svgTD: A Go library for SVG generation #

The library generates SVG as defined by the Scalable Vector Graphics 1.1 Specification (<http://www.w3.org/TR/SVG11/>). 
Output goes to the specified io.Writer.

The inspiration for this library is d3 for syntax and SVGo for readme & examples

## Supported SVG elements and functions ##

### Shapes, lines, text
 circle, ellipse, polygon, polyline, rect (including rounded rects), line, text
 
### Paths
 general, arc, cubic and quadratic bezier paths, 
 
### Transforms ###
 translate, rotate, scale, skewX, skewY

## Building and Usage ##

Usage: (assuming GOPATH is set)

	go get github.com/timdadd/svgTD
	go install github.com/timdadd/svgTD/...
	
	
You can use godoc to browse the documentation from the command line:

	$ go doc github.com/timdadd/svgTD
	

a minimal program, to generate SVG

	package main
	
	import (
		"os"
	)
	
	func main() {
        var width, height = 500, 500
	    svg := svgTD.NewSvgWH(width, height)
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

Drawing in a web server: (http://localhost:2003/circle)

	package main
	
	import (
		"log"
		"net/http"
	)
	
	func main() {
        http.Handle("/circle/", http.HandlerFunc(circle))
	    fmt.Println("http://localhost:2003/circle/")
	    if err := http.ListenAndServe(":2003", nil); err != nil {
		    log.Fatal("ListenAndServe:", err)
	    }
    }

    func circle(w http.ResponseWriter, req *http.Request) {
	    w.Header().Set("Content-Type", "image/svg+xml")
	    svg := svgTD.NewSvgWH(500, 500)
	    svg.Circle().Attrs(250, 250, 125).
		    Style("fill", "none").
		    Style("stroke", "black")
	    _ = svg.Write(w)
    }


### Credits ###

Thanks to svgGO team and D3 for inspiration
