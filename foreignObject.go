package svgTD

// Text places the specified text, t at x,y
//
// Standard Reference: http://www.w3.org/TR/SVG11/text.html#TextElement
type foreignObject struct {
	*Element
}

func (fo *foreignObject) Attr(k string, v ...string) *foreignObject {
	fo.Element.Attr(k, v...)
	return fo
}

// Attrs sets the text attributes, text is optional
func (fo *foreignObject) Attrs(x, y, h, w int) *foreignObject {
	fo.AttrInt("x", x)
	fo.AttrInt("y", y)
	fo.AttrInt("height", h)
	fo.AttrInt("width", w)
	return fo
}

// AttrsFloat sets the text attributes
func (fo *foreignObject) AttrsFloat(x, y float64) *foreignObject {
	fo.AttrFloat("x", x)
	fo.AttrFloat("y", y)
	return fo
}

// AttrsD sets the text attributes with delta x, y
func (fo *foreignObject) AttrsD(dx, dy int) *foreignObject {
	fo.AttrInt("dx", dx)
	fo.AttrInt("dy", dy)
	return fo
}

// AttrsFloatD sets the text attributes with delta x,y
func (fo *foreignObject) AttrsFloatD(dx, dy float64) *foreignObject {
	fo.AttrFloat("dx", dx)
	fo.AttrFloat("dy", dy)
	return fo
}

// Class sets the text class
func (fo *foreignObject) Class(class string) *foreignObject {
	fo.class = class
	return fo
}

// Style sets a style tag on the foreign object
func (fo *foreignObject) Style(k, v string) *foreignObject {
	fo.Element.Style(k, v)
	return fo
}

// Transform sets a transformation on the foreign object
func (fo *foreignObject) Transform(t Transformation) *foreignObject {
	fo.Element.Transform(t)
	return fo
}

// Comment sets a comment on the foreign object
func (fo *foreignObject) Comment(s string) *foreignObject {
	fo.Element.Comment(s)
	return fo
}

// SetText sets the text field, can be multiple fields to build the final text
// If the text field is missing or blank then the text is deleted
func (fo *foreignObject) SetText(tagName string, s ...string) *Element {
	if len(s) == 0 {
		return fo.Element
	}
	var tagElement *Element
	for _, e := range fo.elements {
		if e.tagName == tagName {
			tagElement = e
			break
		}
	}
	if tagElement == nil {
		tagElement = fo.Append(tagName).Attr("xmlns", "http://www.w3.org/1999/xhtml")
	}
	// SetText sets the text field, can be multiple fields to build the final text
	// If the text field is missing or blank then the text is deleted
	tagElement.Attr(tagName, s...)
	//_ = div.SetText(s...)
	return tagElement
}
