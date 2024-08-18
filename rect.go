package svg

// rect is an SVG Element that represents a rectangle. Using this element you can draw rectangles of various
// width, height, with different stroke (outline) and fill colors, with sharp or rounded corners etc.
//
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
type rect struct {
	*Element
}

func (r *rect) Attr(k, v string) *rect {
	r.Element.Attr(k, v)
	return r
}

// Attrs set simple rectangle x, y, width, height
func (r *rect) Attrs(x, y, width, height int) *rect {
	r.AttrInt("x", x)
	r.AttrInt("y", y)
	r.AttrInt("width", width)
	r.AttrInt("height", height)
	return r
}

// FloatAttrs set simple rectangle x, y, width, height
func (r *rect) FloatAttrs(x, y, width, height float64) *rect {
	r.AttrFloat("x", x)
	r.AttrFloat("y", y)
	r.AttrFloat("width", width)
	r.AttrFloat("height", height)
	return r
}

// Center draws a rectangle with its center at x,y, width, height
func (r *rect) Center(x, y, width, height int) *rect {
	return r.CenterFloat(float64(x), float64(y), float64(width), float64(height))
}

// CenterFloat draws a rectangle with its center at x,y, width, height
func (r *rect) CenterFloat(x, y, width, height float64) *rect {
	r.AttrFloat("x", x-(width/2))
	r.AttrFloat("y", y-(height/2))
	r.AttrFloat("width", width)
	r.AttrFloat("height", height)
	return r
}

// Round draws a rounded rectangle with upper the left-hand corner at x,y,width,height.
// The radii for the rounded portion
// are specified by rx (width), and ry (height).
func (r *rect) Round(x, y, width, height, rx, ry int) *rect {
	r.AttrInt("x", x)
	r.AttrInt("y", y)
	r.AttrInt("width", width)
	r.AttrInt("height", height)
	r.AttrInt("rx", rx)
	r.AttrInt("ry", ry)
	return r
}

// RoundFloat draws a rounded rectangle with upper the left-hand corner at x,y,width,height.
// The radii for the rounded portion
// are specified by rx (width), and ry (height).
func (r *rect) RoundFloat(x, y, width, height, rx, ry float64) *rect {
	r.AttrFloat("x", x)
	r.AttrFloat("y", y)
	r.AttrFloat("width", width)
	r.AttrFloat("height", height)
	r.AttrFloat("rx", rx)
	r.AttrFloat("ry", ry)
	return r
}

func (r *rect) Class(class string) *rect {
	r.class = class
	return r
}

func (r *rect) Style(k, v string) *rect {
	r.Element.Style(k, v)
	return r
}

func (r *rect) Transform(t Transformation) *rect {
	r.Element.Transform(t)
	return r
}
