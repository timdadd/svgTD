package svgTD

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Element struct {
	tagName         string
	parent          *Element
	class           string
	attrMap         map[string]string
	attrs           []string          // List of keys into attrMap
	styleMap        map[string]string // list of style settings
	styles          []string          // List of keys into styleMap
	transformations []Transformation
	elements        []*Element
}

func newElement(tagName string, parent *Element) *Element {
	return &Element{
		tagName:         tagName,
		parent:          parent,
		class:           "",
		attrMap:         make(map[string]string, 10),
		attrs:           nil,
		styleMap:        make(map[string]string, 10),
		styles:          nil,
		transformations: nil,
		elements:        nil,
	}
}

// Append adds an element of tagName and returns it e.g. svg.append("line")
// If the append is on <svg> or <g> then append adds as a child otherwise append adds as a peer
func (e *Element) Append(tagName string) (element *Element) {
	switch tagName {
	case "path": // child of current element anyway
		return e.AppendChild(tagName)
	case "marker": // Marker is child of defs and peer of marker
		if e.tagName == "defs" {
			return e.AppendChild(tagName)
		}
	default: // Let element we're appending to decide what happens
		switch e.tagName {
		case "svg", "g", "a": // parent
			return e.AppendChild(tagName)
		}
	}
	return e.AppendPeer(tagName)
}

// AppendChild appends the new element as a child of the current element
// this is an advanced feature, Append is safer
func (e *Element) AppendChild(tagName string) (element *Element) {
	element = newElement(tagName, e)
	e.elements = append(e.elements, element)
	return
}

// AppendPeer appends the new element as a peer of the current element
// this is an advanced feature, Append is safer
func (e *Element) AppendPeer(tagName string) (element *Element) {
	element = newElement(tagName, e.parent)
	e.parent.elements = append(e.parent.elements, element)
	return
}

// Text adds a text element to the parent (e.g. svg.Text)
func (e *Element) Text(s ...string) *text {
	element := e.Append("text")
	element.SetText(s...)
	return &text{element}
}

// TextXY adds a text element to the parent and sets the text X,Y position (e.g. svg.Text)
func (e *Element) TextXY(x, y int, s ...string) *text {
	t := e.Text(s...) // Create the element first
	t.Attrs(x, y)     // Now add x,y
	return t
}

// Line is an SVG Element that draws a straight line between two points.
//
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#LineElement
func (e *Element) Line() *line {
	return &line{e.Append("line")}
}

// Rect is an SVG Element that represents a rectangle. Using this element you can draw rectangles of various
// width, height, with different stroke (outline) and fill colors, with sharp or rounded corners etc.
//
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
func (e *Element) Rect() *rect {
	return &rect{e.Append("rect")}
}

// Circle is an SVG Element that represents a circle. Using this element you can draw circles of various
// radii, with different stroke (outline) and fill colors
//
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#CircleElement
func (e *Element) Circle() *circle {
	return &circle{e.Append("circle")}
}

// Defs is an SVG Element used to store graphical objects that will be used at a later time.
// Objects created inside a Defs element are not rendered directly.
// To display them you have to reference them (with a <use> element for example).
//
// Graphical objects can be referenced from anywhere, however, defining these objects inside a Defs
// element promotes understandability of the SVG content and is beneficial to the overall accessibility of the document.
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
func (e *Element) Defs() *defs {
	return &defs{e.Append("defs")}
}

// SetText sets the text field, can be multiple fields to build the final text
func (e *Element) SetText(s ...string) *Element {
	switch len(s) {
	case 0: // Do nothing
	case 1: // Singleton value
		e.Attr("text", s[0])
	default: // format first
		a := make([]any, len(s)-1)
		for i := range s[1:] {
			a[i] = s[i]
		}
		e.Attr("text", fmt.Sprintf(s[0], a...))
	}
	return e
}

func (e *Element) Class(class string) *Element {
	e.class = class
	return e
}

func (e *Element) Transform(t Transformation) *Element {
	e.transformations = append(e.transformations, t)
	return e
}

// ViewBox controls how much of the item can be viewed
// The value of the viewBox attribute is min-x, min-y, width, and height.
// min-x and min-y represent the smallest X and Y coordinates that the viewBox may have (the origin coordinates of the viewBox)
// and the width and height specify the viewBox size.
func (e *Element) ViewBox(minX, minY, width, height int) *Element {
	e.Attr("viewBox", fmt.Sprintf("%d %d %d %d", minX, minY, width, height))
	return e
}

// writeElement writes out an element either in pretty format or packed format
func writeElement(e *Element, w io.Writer, level int, pretty bool) (err error) {
	if e == nil {
		return
	}
	if pretty {
		if _, err = fmt.Fprint(w, strings.Repeat("  ", level)); err != nil {
			return
		}
	}
	tag := e.Tag()
	openTag := tag.open
	closeTag := tag.close
	if len(e.elements) == 0 && tag.inner == "" {
		openTag = tag.openClose
		closeTag = ""
	}
	if pretty {
		if _, err = fmt.Fprintln(w, openTag); err != nil {
			return
		}
	} else {
		if _, err = fmt.Fprint(w, openTag); err != nil {
			return
		}
	}
	if tag.inner > "" {
		if pretty {
			if _, err = fmt.Fprint(w, strings.Repeat("  ", level+1)); err != nil {
				return
			}
			if err = xml.EscapeText(w, []byte(tag.inner)); err != nil {
				return
			}
			if _, err = fmt.Fprintln(w); err != nil {
				return
			}
		} else {
			if err = xml.EscapeText(w, []byte(tag.inner)); err != nil {
				return
			}
		}
	}
	level++
	for _, elm := range e.elements {
		if err = writeElement(elm, w, level, pretty); err != nil {
			return
		}
	}
	level--

	if pretty {
		if len(e.elements) > 0 || tag.inner > "" {
			if _, err = fmt.Fprintf(w, "%s%s\n", strings.Repeat("  ", level), closeTag); err != nil {
				return
			}
		}
	} else { // packed
		if _, err = fmt.Fprint(w, closeTag); err != nil {
			return
		}
	}
	return
}

// ff formats a floating point number ot leased number of decimals
func ff(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
