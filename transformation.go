package svgTD

import "fmt"

// Transformation makes it possible to transform the shapes created in an SVG image.
// For instance move, scale and rotate the shapes.
// This is a handy way of displaying vertical or diagonal text.
//
// https://www.w3.org/TR/SVG11/coords.html#TransformAttribute
type Transformation interface{}

func (e *Element) transformTags() (s string) {
	if len(e.transformations) == 0 {
		return
	}
	s = ` transform="`
	for _, transformation := range e.transformations {
		switch t := transformation.(type) {
		case translate:
			s += fmt.Sprintf(`translate(%s,%s)`, ff(t.x), ff(t.y))
		case rotate:
			if t.x == 0 && t.y == 0 {
				s += fmt.Sprintf(`rotate(%s)`, ff(t.angle))
			} else {
				s += fmt.Sprintf(`rotate(%s %s %s)`, ff(t.angle), ff(t.x), ff(t.y))
			}
		}
	}
	s += `"`
	return
}

// The translate(x,y) function moves a shape along the x & y axis.
type translate struct {
	x, y float64
}

type rotate struct {
	angle, x, y float64
}

// Translate moves a shape. You pass the x and y value to the translate() function inside the parameters.
// x moves the shape along the x-axis and y moves the shape along the y-axis
func Translate(x, y int) Transformation {
	return translate{float64(x), float64(y)}
}

// TranslateFloat moves a shape. You pass the x and y value to the translate() function inside the parameters.
func TranslateFloat(x, y float64) Transformation {
	return translate{x, y}
}

// Rotate rotates a shape around the point 0,0.
func Rotate(angle int) Transformation {
	return rotate{float64(angle), 0, 0}
}

// RotateFloat rotates a shape around the point 0,0.
func RotateFloat(angle float64) Transformation {
	return rotate{angle, 0, 0}
}

// RotateXY rotates a shape around the point x,y.
func RotateXY(angle, x, y int) Transformation {
	return rotate{float64(angle), float64(x), float64(y)}
}

// RotateFloatXY rotates a shape around the point x,y.
func RotateFloatXY(angle, x, y float64) Transformation {
	return rotate{angle, x, y}
}
