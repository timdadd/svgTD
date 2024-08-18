package svg

import (
	"fmt"
)

// path is an SVG Element that represents the outline of a shape which can be filled, stroked, used as a clipping path,
// or any combination of the three. (See Filling, Stroking and Paint Servers and Clipping, Masking and Compositing.)
//
// A path is described using the concept of a current point. In an analogy with drawing on paper,
// the current point can be thought of as the location of the pen. The position of the pen can be changed,
// and the outline of a shape (open or closed) can be traced by dragging the pen in either straight lines or curves.
//
// Paths represent the geometry of the outline of an object, defined in terms of moveto (set a new current point),
// lineto (draw a straight line), curveto (draw a curve using a cubic Bézier),
// arc (elliptical or circular arc) and closepath (close the current shape by drawing a line to the last moveto)
// elements. Compound paths (i.e., a path with multiple subpaths) are possible to allow effects such as "donut holes"
// in objects.
//
// Reference:https://www.w3.org/TR/SVG11/paths.html#PathElement
type path struct {
	*Element
	d            []pathData
	*pathCommand // Next item can be a path command
}

// Point is s point (x,y)
type point struct {
	path         *path
	x, y         float64
	*point       // Next item can be a point
	*pathCommand // Next item can be a path command
}

// axis is either X or Y
type axis struct {
	path *path
	xy   float64
}

// boolean is either true or false which is 1 or 0 respectively
type boolean struct {
	path *path
	b    bool
}

// pathCommand is a path command
type pathCommand struct {
	path    *path
	command rune
}

// pathData is an item in the path (either a command or point or axis)
type pathData struct {
	pc *pathCommand
	p  *point
	a  *axis
	b  *boolean
}

// arc is a complex ellipse
type arc struct {
	path *path
}

// Path is an SVG Element that represents a graphic path
//
// Reference: https://www.w3.org/TR/SVG11/paths.html#Paths
func (e *Element) Path() *pathCommand {
	return &pathCommand{
		path: &path{
			Element: e.Append("path"),
			d:       []pathData{},
		},
	}
}

// pathDataTag returns the value of the "d" attr
func (p *path) append(pd pathData) {
	var pathText string
	var d = p.Element.attrMap["d"]
	if d > "" {
		d += " "
	}
	if pd.pc != nil {
		pathText = string(pd.pc.command)
	} else {
		if pd.p != nil {
			pathText = fmt.Sprintf("%s %s", ff(pd.p.x), ff(pd.p.y))
		} else if pd.a != nil {
			pathText = ff(pd.a.xy)
		} else if pd.b != nil {
			if pd.b.b {
				pathText = "1"
			} else {
				pathText = "0"
			}
		}
	}
	p.d = append(p.d, pd)
	p.Element.Attr("d", d+pathText)
}

// MoveTo moves to a Relative or Absolute x,y position
// The MoveTo command establishes a new current point.
// The effect is as if the "pen" were lifted and moved to a new location.
func (pc *pathCommand) MoveTo(absolute bool) *point {
	if absolute {
		pc.command = 'M'
	} else {
		pc.command = 'm'
	}
	pc.path.append(pathData{pc: pc})
	return &point{path: pc.path, pathCommand: &pathCommand{path: pc.path}}
}

// LineTo draws a line to a Relative or absolute x,y position
// The LineTo command establishes a new current point.
// The effect is as if the "pen" was drawing a line to a new location.
func (pc *pathCommand) LineTo(absolute bool) *point {
	if absolute {
		pc.command = 'L'
	} else {
		pc.command = 'l'
	}
	pc.path.append(pathData{pc: pc})
	return &point{path: pc.path, pathCommand: &pathCommand{path: pc.path}}
}

// CurveTo draws a curve or curves to a new Relative or Absolute x,y position
// Draws a cubic Bézier curve from the current point to (x,y) using (x1,y1) as the control point at the beginning
// of the curve and (x2,y2) as the control point at the end of the curve.
// Multiple sets of coordinates may be specified to draw a polybézier. At the end, the new current point becomes
// the final (x,y) coordinate pair used in the polybézier
func (pc *pathCommand) CurveTo(absolute bool) *point {
	if absolute {
		pc.command = 'C'
	} else {
		pc.command = 'c'
	}
	pc.path.append(pathData{pc: pc})
	return &point{path: pc.path, pathCommand: &pathCommand{path: pc.path}}
}

// P is an integer point
func (p *point) P(x, y int) *point {
	return p.Pf(float64(x), float64(y))
}

// Pf is a floating point
func (p *point) Pf(x, y float64) *point {
	p.x, p.y = x, y
	p.path.append(pathData{p: p})
	return &point{path: p.path, pathCommand: &pathCommand{path: p.path}}
}

// HorizontalLineTo draws a horizontal line to a new Relative or absolute x position
// The HorizontalLineTo command establishes a new current point.
// The effect is as if the "pen" was drawing a horizontal line to a new location.
func (pc *pathCommand) HorizontalLineTo(absolute bool) *axis {
	if absolute {
		pc.command = 'H'
	} else {
		pc.command = 'h'
	}
	pc.path.append(pathData{pc: pc})
	return &axis{path: pc.path}
}

// VerticalLineTo draws a Vertical line to a new Relative or absolute y position
// The VerticalLineTo command establishes a new current point.
// The effect is as if the "pen" was drawing a Vertical line to a new location.
func (pc *pathCommand) VerticalLineTo(absolute bool) *axis {
	if absolute {
		pc.command = 'V'
	} else {
		pc.command = 'v'
	}
	pc.path.append(pathData{pc: pc})
	return &axis{path: pc.path}
}

// A is an integer Axis either X or Y
func (a *axis) A(xy int) *pathCommand {
	return a.Af(float64(xy))
}

// Af is a floating axis either X or Y
func (a *axis) Af(xy float64) *pathCommand {
	a.xy = xy
	a.path.append(pathData{a: a})
	return &pathCommand{path: a.path}
}

// Arc draws an elliptical arc, with optional style, beginning coordinate at current point (x,y)
// ending coordinate at ex, ey
// width and height of the arc are specified by ax, ay, the x axis rotation is r
// if sweep is true, then the arc will be drawn in a "positive-angle" direction (clockwise), if false,
// the arc is drawn counterclockwise.
// if large is true, the arc sweep angle is greater than or equal to 180 degrees,
// otherwise the arc sweep is less than 180 degrees
// http://www.w3.org/TR/SVG11/paths.html#PathDataEllipticalArcCommands
func (pc *pathCommand) Arc(absolute bool) *arc {
	if absolute {
		pc.command = 'A'
	} else {
		pc.command = 'a'
	}
	pc.path.append(pathData{pc: pc})
	return &arc{path: pc.path}
}

// Attrs sets the attrs of the arc
func (a *arc) Attrs(ax, ay, r int, large, sweep bool, ex, ey int) *pathCommand {
	a.path.append(pathData{p: &point{x: float64(ax), y: float64(ay)}})
	a.path.append(pathData{a: &axis{xy: float64(r)}})
	a.path.append(pathData{b: &boolean{b: large}})
	a.path.append(pathData{b: &boolean{b: sweep}})
	a.path.append(pathData{p: &point{x: float64(ex), y: float64(ey)}})
	return &pathCommand{path: a.path}
}

func (pc *pathCommand) Element() *Element {
	return pc.path.Element
}
