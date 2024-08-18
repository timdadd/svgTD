package svg

// Text places the specified text, t at x,y
//
// Standard Reference: http://www.w3.org/TR/SVG11/text.html#TextElement
type text struct {
	*Element
}

func (txt *text) Attr(k string, v ...string) *text {
	txt.Element.Attr(k, v...)
	return txt
}

// Attrs sets the text attributes, text is optional
func (txt *text) Attrs(x, y int, s ...string) *text {
	txt.AttrInt("x", x)
	txt.AttrInt("y", y)
	if len(s) > 0 {
		txt.SetText(s...)
	}
	return txt
}

// AttrsFloat sets the text attributes
func (txt *text) AttrsFloat(x, y float64, s ...string) *text {
	txt.AttrFloat("x", x)
	txt.AttrFloat("y", y)
	if len(s) > 0 {
		txt.SetText(s...)
	}
	return txt.SetText(s...)
}

// AttrsD sets the text attributes with delta x, y
func (txt *text) AttrsD(dx, dy int, s ...string) *text {
	txt.AttrInt("dx", dx)
	txt.AttrInt("dy", dy)
	return txt
}

// AttrsFloatD sets the text attributes with delta x,y
func (txt *text) AttrsFloatD(dx, dy float64, s ...string) *text {
	txt.AttrFloat("dx", dx)
	txt.AttrFloat("dy", dy)
	return txt.SetText(s...)
}

// Class sets the text class
func (txt *text) Class(class string) *text {
	txt.class = class
	return txt
}

// Style sets a style tag on the text
func (txt *text) Style(k, v string) *text {
	txt.Element.Style(k, v)
	return txt
}

// Transform sets a transformation on the text
func (txt *text) Transform(t Transformation) *text {
	txt.Element.Transform(t)
	return txt
}

// SetText sets the text field, can be multiple fields to build the final text
// If the text field is missing or blank then the text is deleted
func (txt *text) SetText(s ...string) *text {
	txt.Attr("text", s...)
	return txt
}
