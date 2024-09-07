package svgTD

import "fmt"

type Tag struct {
	open      string
	inner     string
	close     string
	openClose string
}

// Tag returns the opening and closing tag for this element
func (e *Element) Tag() (tag Tag) {
	tag.close = fmt.Sprintf("</%s>", e.tagName)
	tag.open = fmt.Sprintf("<%s", e.tagName)
	if e.class > "" {
		tag.open += fmt.Sprintf(` class="%s"`, e.class)
	}
	// Handle "text" attr
	if t, inMap := e.attrMap["text"]; inMap {
		tag.inner = t
	}
	tag.open += e.attrTags()
	tag.open += e.styleTag()
	tag.open += e.transformTags()
	tag.openClose = tag.open + "/>"
	tag.open += ">"
	return
}
