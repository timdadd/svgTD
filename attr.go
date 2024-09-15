package svgTD

import (
	"fmt"
	"strconv"
)

// Attr sets an attr on the element
// If no value is sent then the attr is deleted
func (e *Element) Attr(k string, v ...string) *Element {
	if len(v) == 0 || v[0] == "" {
		delete(e.attrMap, k)
		return e
	}
	if _, ok := e.attrMap[k]; !ok {
		e.attrs = append(e.attrs, k)
	}
	if len(v) == 1 {
		e.attrMap[k] = v[0]
	} else {
		// If multiple values then assume this is format, values
		var anyList = make([]any, len(v)-1)
		for i, s := range v[1:] {
			anyList[i] = any(s)
		}
		e.attrMap[k] = fmt.Sprintf(v[0], anyList...)
	}
	return e
}

// AttrInt sets a integer attr on the element
func (e *Element) AttrInt(k string, v int) *Element {
	return e.Attr(k, strconv.Itoa(v))
}

// AttrFloat sets a floating point attr on the element
func (e *Element) AttrFloat(k string, v float64) *Element {
	return e.Attr(k, ff(v))
}

// attrTags returns all the Attr tags in the order they were set
func (e *Element) attrTags() (s string) {
	// attrs are put in the order they were set
	for _, k := range e.attrs {
		if k == e.tagName {
			continue
		}
		//if k == "text" && e.tagName == "text" {
		//	continue
		//}
		if v, inMap := e.attrMap[k]; inMap {
			s += fmt.Sprintf(` %s="%s"`, k, v) // Quotes with attrs
		}
	}
	return s
}

// braceTags returns all the Attr tags in the order they were set but line by line for
func (e *Element) braceTags() (s []string) {
	// attrs are put in the order they were set
	var braceAttrs []string
	for _, k := range e.attrs {
		if k == e.tagName {
			continue
		}
		if v, inMap := e.attrMap[k]; inMap {
			braceAttrs = append(braceAttrs, fmt.Sprintf(`%s: %s;`, k, v))
		}
	}
	return braceAttrs
}
