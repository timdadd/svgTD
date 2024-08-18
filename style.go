package svg

import (
	"fmt"
	"strconv"
)

// Style sets a Style on the element
func (e *Element) Style(k, v string) *Element {
	if _, ok := e.styleMap[k]; !ok {
		e.styles = append(e.styles, k)
	}
	e.styleMap[k] = v
	return e
}

// StyleInt sets a integer Style on the element
func (e *Element) StyleInt(k string, v int) *Element {
	return e.Style(k, strconv.Itoa(v))
}

// StyleFloat sets a floating point attr on the element
func (e *Element) StyleFloat(k string, v float64) *Element {
	return e.Style(k, ff(v))
}

// styleTag returns all the Style Attrs in the order they were set
func (e *Element) styleTag() (s string) {
	for _, k := range e.styles {
		if v, inMap := e.styleMap[k]; inMap {
			if len(s) == 0 {
				s = ` style="`
			}
			s += fmt.Sprintf("%s:%s;", k, v) // No quotes in style
		}
	}
	if len(s) > 0 {
		s += "\""
	}
	return
}
