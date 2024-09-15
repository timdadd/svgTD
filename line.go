package svgTD

// Line is an SVG Element that draws a straight line between two points.
//
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#LineElement
type line struct {
	*Element
}

// Attrs sets the line From To
func (l *line) Attrs(x1, y1, x2, y2 int) *line {
	l.AttrInt("x1", x1)
	l.AttrInt("y1", y1)
	l.AttrInt("x2", x2)
	l.AttrInt("y2", y2)
	return l
}

// AttrsFloat sets the line attributes
func (l *line) AttrsFloat(x1, y1, x2, y2 float64) *line {
	l.AttrFloat("x1", x1)
	l.AttrFloat("y1", y1)
	l.AttrFloat("x2", x2)
	l.AttrFloat("y2", y2)
	return l
}

func (l *line) Class(class string) *line {
	l.Element.Class(class)
	return l
}

func (l *line) Style(k, v string) *line {
	l.Element.Style(k, v)
	return l
}

func (l *line) Transform(t Transformation) *line {
	l.Element.Transform(t)
	return l
}
func (l *line) Comment(c string) *line {
	l.Element.Comment(c)
	return l
}
