package svgTD

// Defs is an SVG Element used to store graphical objects that will be used at a later time.
// Objects created inside a Defs element are not rendered directly.
// To display them you have to reference them (with a <use> element for example).
//
// Graphical objects can be referenced from anywhere, however, defining these objects inside a Defs
// element promotes understandability of the SVG content and is beneficial to the overall accessibility of the document.
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
type defs struct {
	*Element
}
