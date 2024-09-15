package svgTD

import "fmt"

// Polygon defines a closed shape consisting of a set of connected straight line segments.
// The last point is connected to the first point.
//
// Standard Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polygon
type polygon struct {
	*Element
	points []polygonPoint
}

// polygonPoint is s point (x,y)
type polygonPoint struct {
	polygon       *polygon
	x, y          float64
	*polygonPoint // Next item can be a point
}

func (e *Element) Polygon() *polygonPoint {
	return &polygonPoint{
		polygon: &polygon{
			Element: e.Append("polygon"),
			points:  []polygonPoint{},
		},
	}
}

// append updates the "points" att
func (p *polygon) append(pt *polygonPoint) {
	var points = p.Element.attrMap["points"]
	if points > "" {
		points += " "
	}
	if p != nil {
		points += fmt.Sprintf("%s,%s", ff(pt.x), ff(pt.y))
		p.Element.Attr("points", points)
	}
	p.points = append(p.points, *pt)
}

// Points is a list of points
func (pp *polygonPoint) Points(xy ...int) *polygonPoint {
	for i := range xy {
		if i%2 == 0 && i < len(xy)-1 {
			pp = pp.Pf(float64(xy[i]), float64(xy[i+1]))
		}
	}
	return pp
}

// P is an integer point
func (pp *polygonPoint) P(x, y int) *polygonPoint {
	return pp.Pf(float64(x), float64(y))
}

// Pf is a floating point
func (pp *polygonPoint) Pf(x, y float64) *polygonPoint {
	pp.x, pp.y = x, y
	pp.polygon.append(pp)
	return &polygonPoint{polygon: pp.polygon}
}

func (pp *polygonPoint) Class(class string) *polygonPoint {
	pp.polygon.Element.Class(class)
	return pp
}

func (pp *polygonPoint) Style(k, v string) *polygonPoint {
	pp.polygon.Element.Style(k, v)
	return pp
}

func (pp *polygonPoint) Transform(t Transformation) *polygonPoint {
	pp.polygon.Element.Transform(t)
	return pp
}
func (pp *polygonPoint) Comment(c string) *polygonPoint {
	pp.polygon.Element.Comment(c)
	return pp
}
