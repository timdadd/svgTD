package sequenceDiagram_test

import (
	"fmt"
	"svgTD"
	"testing"
)

const (
	expectedPrettySD = `<svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
    <style>
        div {
            color: white;
            font: 18px serif;
            height: 100%;
            overflow: auto;
        }
    </style>
    <polygon points="5,5 195,10 185,185 10,195" />
    <!-- Common use case: embed HTML text into SVG -->
    <foreignObject x="20" y="20" width="160" height="160">
        <!--
        In the context of SVG embedded in an HTML document, the XHTML
        namespace could be omitted, but it is mandatory in the
        context of an SVG document
        -->
        <div xmlns="http://www.w3.org/1999/xhtml">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
            mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
            imperdiet eros. Aliquam erat volutpat.
        </div>
    </foreignObject>
</svg>
`
	expectedPackedSD = `<svg xmlns="http://www.w3.org/2000/svg" width="800" height="600"><line x1="50" y1="20" x2="50" y2="480" style="stroke:#888;"/><line x1="150" y1="20" x2="150" y2="480" style="stroke:#888;"/><line x1="250" y1="20" x2="250" y2="480" style="stroke:#888;"/><line x1="350" y1="20" x2="350" y2="480" style="stroke:#888;"/><line x1="450" y1="20" x2="450" y2="480" style="stroke:#888;"/><line x1="550" y1="20" x2="550" y2="480" style="stroke:#888;"/><line x1="650" y1="20" x2="650" y2="480" style="stroke:#888;"/><g class="systems" transform="translate(50,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(150,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(250,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(350,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(450,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(550,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="systems" transform="translate(650,20)"><rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"/></g><g class="labels" transform="translate(50,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class A</text></g><g class="labels" transform="translate(150,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class B</text></g><g class="labels" transform="translate(250,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class C</text></g><g class="labels" transform="translate(350,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class D</text></g><g class="labels" transform="translate(450,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class E</text></g><g class="labels" transform="translate(550,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class F</text></g><g class="labels" transform="translate(650,20)"><text dx="0" dy="25" style="text-anchor:middle;">Class G</text></g><line x1="50" y1="100" x2="250" y2="100" marker-end="url(#end)" style="stroke:black;"/><text>From A to C</text><line x1="250" y1="150" x2="350" y2="150" marker-end="url(#end)" style="stroke:black;"/><text>From C to D</text><line x1="350" y1="200" x2="450" y2="200" marker-end="url(#end)" style="stroke:black;"/><text>From D to E</text><line x1="450" y1="250" x2="350" y2="250" marker-end="url(#end)" style="stroke:black;"/><text>From E to D</text><line x1="350" y1="300" x2="650" y2="300" marker-end="url(#end)" style="stroke:black;"/><text>From D to G</text><line x1="650" y1="350" x2="350" y2="350" marker-end="url(#end)" style="stroke:black;"/><text>From G to D</text><line x1="350" y1="400" x2="250" y2="400" marker-end="url(#end)" style="stroke:black;"/><text>From D to C</text><line x1="250" y1="450" x2="50" y2="450" marker-end="url(#end)" style="stroke:black;"/><text>From C to A</text><g class="From A to C" transform="translate(150,90)"><text style="text-anchor:middle;">From A to C</text></g><g class="From C to D" transform="translate(300,140)"><text style="text-anchor:middle;">From C to D</text></g><g class="From D to E" transform="translate(400,190)"><text style="text-anchor:middle;">From D to E</text></g><g class="From E to D" transform="translate(400,240)"><text style="text-anchor:middle;">From E to D</text></g><g class="From D to G" transform="translate(500,290)"><text style="text-anchor:middle;">From D to G</text></g><g class="From G to D" transform="translate(500,340)"><text style="text-anchor:middle;">From G to D</text></g><g class="From D to C" transform="translate(300,390)"><text style="text-anchor:middle;">From D to C</text></g><g class="From C to A" transform="translate(150,440)"><text style="text-anchor:middle;">From C to A</text></g><defs><marker id="end" viewBox="0 -5 10 10" refX="10" refY="0" markerWidth="10" markerHeight="10" orient="auto"><path d="M 0 -5 L 10 0 0 5"/></marker></defs></svg>`
)

// TestForeignObject is a GO Implementation of https://developer.mozilla.org/en-US/docs/Web/SVG/Element/foreignObject
func TestForeignObject(t *testing.T) {
	// Create the s canvas
	var s = svgTD.New().ViewBox(0, 0, 200, 200)
	s.Element.Append("style").Append("div").
		Attr("color", "white").
		Attr("font", "18px serif").
		Attr("height", "100%").
		Attr("overflow", "auto")
	s.Polygon().Points(5, 5, 195, 10, 185, 185, 10, 195).
		Comment("Common use case: embed HTML text into SVG")
	s.Text("hello")
	s.ForeignObjectXYHW(20, 20, 160, 160).
		SetText("div",
			`            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis mollis
            mi ut ultricies. Nullam magna ipsum, porta vel dui convallis, rutrum
            imperdiet eros. Aliquam erat volutpat.`)
	s.Pretty(true)
	if b, err := s.Marshal(); err != nil {
		t.Errorf("cannot generate SVG %v", err)
	} else {
		fmt.Println(string(b))
	}
	//else {
	//	if string(b) != expectedPrettySD {
	//		if len(b) != len(expectedPrettySD) {
	//			t.Errorf("length error, expected %d, received %d", len(expectedPrettySD), len(b))
	//		}
	//		for i, a := range string(b) {
	//			fmt.Printf("%d. %s %s\n", i, string(expectedPrettySD[i:i+1]), string(b[i]))
	//			if expectedPrettySD[i:i+1] != string(a) {
	//				println(string(b))
	//				t.Errorf("%f2%% good", float64(i*100)/float64(len(expectedPrettySD)))
	//				return
	//			}
	//		}
	//	}
	//	fmt.Println("Pretty all good")
	//}
	//
	//s.Pretty(false)
	//if b, err := s.Marshal(); err != nil {
	//	t.Errorf("cannot generate SVG %v", err)
	//} else {
	//	if string(b) != expectedPackedSD {
	//		if len(b) != len(expectedPackedSD) {
	//			t.Errorf("length error, expected %d, received %d", len(expectedPackedSD), len(b))
	//		}
	//		for i, a := range string(b) {
	//			fmt.Printf("%d. %s %s\n", i, string(expectedPackedSD[i:i+1]), string(b[i]))
	//			if expectedPackedSD[i:i+1] != string(a) {
	//				println(string(b))
	//				t.Errorf("%f2%% good", float64(i*100)/float64(len(expectedPackedSD)))
	//				return
	//			}
	//		}
	//	}
	//	fmt.Println("Packed All good")
	//}
	return

}
