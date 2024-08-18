package svg

// marker is an SVG Element that represents a graphic used for drawing arrowheads or polymarkers on a
// given <path>, <line>, <polyline> or <polygon> element.
//
// # Markers can be attached to shapes using the marker-start, marker-mid, and marker-end properties
//
// Standard Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
type marker struct {
	*Element
}

// Marker is an SVG Element used to create a marker for the start, mid, and end of a <line>, <path>, <polyline> or <polygon>.
//
// All SVG markers are defined within a <defs> element. The <defs> element is short for "definitions",
// and contains definition of special elements (such as markers).
//
// The marker is attached to the shapes using the marker-start, marker-mid, and marker-end attributes.
//
// The <marker> element has six basic attributes to position and set the size of the marker
func (d *defs) Marker(id string) *marker {
	element := d.AppendChild("marker")
	element.Attr("id", id)
	return &marker{element}
}

// Marker is an SVG Element that represents a graphic used for drawing arrowheads or polymarkers on a
// given <path>, <line>, <polyline> or <polygon> element.
//
// # Markers can be attached to shapes using the marker-start, marker-mid, and marker-end properties
//
// Standard Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
func (m *marker) Marker(id string) *marker {
	element := m.AppendPeer("marker")
	element.Attr("id", id)
	return &marker{element}
}

// ViewBox controls how much of the item can be viewed
// The value of the viewBox attribute is min-x, min-y, width, and height.
// min-x and min-y represent the smallest X and Y coordinates that the viewBox may have (the origin coordinates of the viewBox)
// and the width and height specify the viewBox size.
func (m *marker) ViewBox(minX, minY, width, height int) *marker {
	m.Element.ViewBox(minX, minY, width, height)
	return m
}

// Ref defines the x & y coordinate for the reference point of the marker
// Value type: left|center|right|<coordinate>
func (m *marker) Ref(x, y string) *marker {
	m.Attr("refX", x)
	m.Attr("refY", y)
	return m
}
