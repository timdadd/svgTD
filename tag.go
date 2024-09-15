package svgTD

import "fmt"

type Tag struct {
	open        string
	escapeInner bool
	inner       []string
	close       string
	openClose   string
}

// Tag returns the opening and closing tag for this element
func (e *Element) Tag() (tag Tag) {
	if e.parent != nil && e.parent.tagName == "style" {
		tag.open = fmt.Sprintf("%s {", e.tagName)
		tag.inner = e.braceTags()
		tag.close = "}"
		return
	}
	tag.close = fmt.Sprintf("</%s>", e.tagName)
	tag.open = fmt.Sprintf("<%s", e.tagName)
	if e.class > "" {
		tag.open += fmt.Sprintf(` class="%s"`, e.class)
	}
	// Handle "text" attr
	//if t, inMap := e.attrMap["text"]; inMap {
	// If the tag name and element name are the same then make inner
	if t, inMap := e.attrMap[e.tagName]; inMap {
		tag.inner = []string{t}
		tag.escapeInner = true
	}
	tag.open += e.attrTags()
	tag.open += e.styleTag()
	tag.open += e.transformTags()
	tag.openClose = tag.open + "/>"
	tag.open += ">"
	return
}
