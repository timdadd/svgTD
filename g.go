package svg

// The SVG <g> element is used to group SVG shapes together. Once grouped you can transform the whole group
// of shapes as if it was a single shape. This is an advantage compared to a nested <svg> element which cannot
// be the target of transformation by itself.
//
// You can also style the grouped elements, and reuse them as if they were a single element.
//
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#GElement
type group struct {
	*Element
}

// G creates a new group as child of SVG
func (svg *SVG) G() *group {
	g := &group{Element: newElement("g", svg.Element)}
	svg.elements = append(svg.elements, g.Element)
	return g
}

// SVG re-homes an SVG onto this group
// We should check the SVG isn't a parent of this group otherwise we get infinite recursion
func (g *group) SVG(s *SVG) *group {
	g.Element.elements = append(g.elements, s.Element)
	return g
}

func (g *group) Style(k, v string) *group {
	g.Element.Style(k, v)
	return g
}

func (g *group) Class(class string) *group {
	g.Element.Class(class)
	return g
}

func (g *group) Transform(t Transformation) *group {
	g.Element.Transform(t)
	return g
}

//func (g *group) ToString() string {
//	return fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" %s`, l.x1, l.y1, l.x2, l.y2, endStyle(l.styles, emptyclose))
//}

//// Gstyle begins a group, with the specified style.
//// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#GElement
//func (svg *SVG) Gstyle(s string) { svg.println(group("style", s)) }

//// Gtransform begins a group, with the specified transform
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) Gtransform(s string) { svg.println(group("transform", s)) }
//
//// Translate begins coordinate translation, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) Translate(x, y int) { svg.Gtransform(translate(x, y)) }
//
//// Scale scales the coordinate system by n, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) Scale(n float64) { svg.Gtransform(scale(n)) }
//
//// ScaleXY scales the coordinate system by dx and dy, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) ScaleXY(dx, dy float64) { svg.Gtransform(scaleXY(dx, dy)) }
//
//// SkewX skews the x coordinate system by angle a, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) SkewX(a float64) { svg.Gtransform(skewX(a)) }
//
//// SkewY skews the y coordinate system by angle a, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) SkewY(a float64) { svg.Gtransform(skewY(a)) }
//
//// SkewXY skews x and y coordinates by ax, ay respectively, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) SkewXY(ax, ay float64) { svg.Gtransform(skewX(ax) + " " + skewY(ay)) }
//
//// Rotate rotates the coordinate system by r degrees, end with Gend()
//// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
//func (svg *SVG) Rotate(r float64) { svg.Gtransform(rotate(r)) }
//
//// TranslateRotate translates the coordinate system to (x,y), then rotates to r degrees, end with Gend()
//func (svg *SVG) TranslateRotate(x, y int, r float64) {
//	svg.Gtransform(translate(x, y) + " " + rotate(r))
//}
//
//// RotateTranslate rotates the coordinate system r degrees, then translates to (x,y), end with Gend()
//func (svg *SVG) RotateTranslate(x, y int, r float64) {
//	svg.Gtransform(rotate(r) + " " + translate(x, y))
//}
//
//// Group begins a group with arbitrary attributes
//func (svg *SVG) Group(s ...string) { svg.printf("<g %s\n", endstyle(s, `>`)) }
//
//// Gid begins a group, with the specified id
//func (svg *SVG) Gid(s string) {
//	svg.print(`<g id="`)
//	xml.Escape(svg.Writer, []byte(s))
//	svg.println(`">`)
//}
//
//// Gend ends a group (must be paired with Gsttyle, Gtransform, Gid).
//func (svg *SVG) Gend() { svg.println(`</g>`) }
