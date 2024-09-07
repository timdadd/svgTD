package svgTD

// circle is an SVG Element that represents a circle. Using this element you can draw circles of various
// sizes (radius), with different stroke (outline) and fill colors
// circle centered at x,y, with radius r
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#CircleElement
type circle struct {
	*Element
}

func (c *circle) Attr(k, v string) *circle {
	c.Element.Attr(k, v)
	return c
}

// Attrs set simple circle with center at x, y of given radius
func (c *circle) Attrs(cx, cy, r int) *circle {
	c.AttrInt("cx", cx)
	c.AttrInt("cy", cy)
	c.AttrInt("r", r)
	return c
}

// FloatAttrs set simple circle with center at x, y of given radius
func (c *circle) FloatAttrs(cx, cy, r float64) *circle {
	c.AttrFloat("cx", cx)
	c.AttrFloat("cy", cy)
	c.AttrFloat("r", r)
	return c
}

func (c *circle) Class(class string) *circle {
	c.class = class
	return c
}

func (c *circle) Style(k, v string) *circle {
	c.Element.Style(k, v)
	return c
}

func (c *circle) Transform(t Transformation) *circle {
	c.Element.Transform(t)
	return c
}
