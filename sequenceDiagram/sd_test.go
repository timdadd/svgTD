package sequenceDiagram_test

import (
	"fmt"
	"svgTD"
	"testing"
)

const expectedSD = `<svg xmlns="http://www.w3.org/2000/svg" width="800" height="600">
  <line x1="50" y1="20" x2="50" y2="480" style="stroke:#888;"></line>
  <line x1="150" y1="20" x2="150" y2="480" style="stroke:#888;"></line>
  <line x1="250" y1="20" x2="250" y2="480" style="stroke:#888;"></line>
  <line x1="350" y1="20" x2="350" y2="480" style="stroke:#888;"></line>
  <line x1="450" y1="20" x2="450" y2="480" style="stroke:#888;"></line>
  <line x1="550" y1="20" x2="550" y2="480" style="stroke:#888;"></line>
  <line x1="650" y1="20" x2="650" y2="480" style="stroke:#888;"></line>
  <g class="systems" transform="translate(50,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(150,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(250,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(350,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(450,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(550,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="systems" transform="translate(650,20)">
    <rect x="-40" y="0" width="80" height="40" style="fill:#CCC;"></rect>
  </g>
  <g class="labels" transform="translate(50,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class A
    </text>
  </g>
  <g class="labels" transform="translate(150,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class B
    </text>
  </g>
  <g class="labels" transform="translate(250,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class C
    </text>
  </g>
  <g class="labels" transform="translate(350,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class D
    </text>
  </g>
  <g class="labels" transform="translate(450,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class E
    </text>
  </g>
  <g class="labels" transform="translate(550,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class F
    </text>
  </g>
  <g class="labels" transform="translate(650,20)">
    <text dx="0" dy="25" style="text-anchor:middle;">
      Class G
    </text>
  </g>
  <line x1="50" y1="100" x2="250" y2="100" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From A to C
  </text>
  <line x1="250" y1="150" x2="350" y2="150" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From C to D
  </text>
  <line x1="350" y1="200" x2="450" y2="200" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From D to E
  </text>
  <line x1="450" y1="250" x2="350" y2="250" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From E to D
  </text>
  <line x1="350" y1="300" x2="650" y2="300" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From D to G
  </text>
  <line x1="650" y1="350" x2="350" y2="350" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From G to D
  </text>
  <line x1="350" y1="400" x2="250" y2="400" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From D to C
  </text>
  <line x1="250" y1="450" x2="50" y2="450" marker-end="url(#end)" style="stroke:black;"></line>
  <text>
    From C to A
  </text>
  <g class="From A to C" transform="translate(150,90)">
    <text style="text-anchor:middle;">
      From A to C
    </text>
  </g>
  <g class="From C to D" transform="translate(300,140)">
    <text style="text-anchor:middle;">
      From C to D
    </text>
  </g>
  <g class="From D to E" transform="translate(400,190)">
    <text style="text-anchor:middle;">
      From D to E
    </text>
  </g>
  <g class="From E to D" transform="translate(400,240)">
    <text style="text-anchor:middle;">
      From E to D
    </text>
  </g>
  <g class="From D to G" transform="translate(500,290)">
    <text style="text-anchor:middle;">
      From D to G
    </text>
  </g>
  <g class="From G to D" transform="translate(500,340)">
    <text style="text-anchor:middle;">
      From G to D
    </text>
  </g>
  <g class="From D to C" transform="translate(300,390)">
    <text style="text-anchor:middle;">
      From D to C
    </text>
  </g>
  <g class="From C to A" transform="translate(150,440)">
    <text style="text-anchor:middle;">
      From C to A
    </text>
  </g>
  <defs>
    <marker id="end" viewBox="0 -5 10 10" refX="10" refY="0" markerWidth="10" markerHeight="10" orient="auto">
      <path d="M 0 -5 L 10 0 0 5"></path>
    </marker>
  </defs>
</svg>
`

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

// TestSequenceDiagram is a GO implementation of joegaffey/D3 sequence diagram
func TestSequenceDiagram(t *testing.T) {
	var classes = []string{"Class A", "Class B", "Class C", "Class D", "Class E", "Class F", "Class G"}

	type message struct {
		start   int
		end     int
		message string
	}
	var messages = []message{{start: 0, end: 2, message: "From A to C"},
		{start: 2, end: 3, message: "From C to D"},
		{start: 3, end: 4, message: "From D to E"},
		{start: 4, end: 3, message: "From E to D"},
		{start: 3, end: 6, message: "From D to G"},
		{start: 6, end: 3, message: "From G to D"},
		{start: 3, end: 2, message: "From D to C"},
		{start: 2, end: 0, message: "From C to A"}}

	var xPadding = 50
	var yPadding = 20
	var verticalSpacing = 100
	var verticalPadding = 60

	var classWidth = 80
	var classHeight = 40
	var classLavelOffsetY = 25

	var messageSpace = 50
	var messageLabelOffsetY = 70
	var messageArrowOffsetY = 80

	var canvasWidth = 800
	var canvasHeight = 600

	// Create the s canvas
	var s = svg.New().
		Width(canvasWidth).
		Height(canvasHeight)

	// Draw vertical lines
	for i := range classes {
		s.Line().
			Attrs(xPadding+i*verticalSpacing, yPadding, xPadding+i*verticalSpacing, yPadding+verticalPadding+len(messages)*messageSpace).
			Style("stroke", "#888")
	}

	// Draw classes
	for i := range classes {
		x := xPadding + i*verticalSpacing
		s.G().Transform(svg.Translate(x, yPadding)).
			Class("systems").
			Rect().
			Attrs(-classWidth/2, 0, classWidth, classHeight).
			Style("fill", "#CCC")
	}

	// Draw class labels in the class boxes
	for i, c := range classes {
		var x = xPadding + i*verticalSpacing
		s.G().
			Transform(svg.Translate(x, yPadding)).
			Class("labels").
			Text(c).Style("text-anchor", "middle").
			AttrsD(0, classLavelOffsetY)
	}

	// Draw the message arrows
	for i, m := range messages {
		var y = yPadding + messageArrowOffsetY + i*messageSpace
		s.Line().Style("stroke", "black").
			Attrs(xPadding+m.start*verticalSpacing, y, xPadding+m.end*verticalSpacing, y).
			Attr("marker-end", "url(#end)").
			Text(m.message)
	}

	// Draw Message Labels
	for i, m := range messages {
		var xPos = xPadding + (((m.end - m.start) * verticalSpacing) / 2) + (m.start * verticalSpacing)
		var yPos = yPadding + messageLabelOffsetY + i*messageSpace
		s.G().Transform(svg.Translate(xPos, yPos)).
			Class(m.message).
			Text(m.message).Style("text-anchor", "middle")
	}

	// Arrow Style
	s.Defs().Marker("end").
		ViewBox(0, -5, 10, 10).
		Ref("10", "0").
		AttrInt("markerWidth", 10).
		AttrInt("markerHeight", 10).
		Attr("orient", "auto").
		Path().
		MoveTo(true).P(0, -5).
		LineTo(true).P(10, 0).P(0, 5)

	s.Pretty(true)
	if b, err := s.Marshal(); err != nil {
		t.Errorf("cannot generate SVG %v", err)
	} else {
		if string(b) != expectedSD {
			if len(b) != len(expectedSD) {
				t.Errorf("length error, expected %d, received %d", len(expectedSD), len(b))
			}
			for i, a := range string(b) {
				fmt.Printf("%d. %s %s\n", i, string(expectedSD[i:i+1]), string(b[i]))
				if expectedSD[i:i+1] != string(a) {
					println(string(b))
					t.Errorf("%f2%% good", float64(i*100)/float64(len(expectedSD)))
					return
				}
			}
		}
		fmt.Println("All good")
	}
	return

}
