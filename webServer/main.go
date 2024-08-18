package main

import (
	"fmt"
	"log"
	"net/http"
	"svgTD"
)

func main() {
	http.Handle("/circle/", http.HandlerFunc(circle))
	fmt.Println("http://localhost:2003/circle/")
	if err := http.ListenAndServe(":2003", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	canvas := svg.NewWH(500, 500)
	canvas.Circle().Attrs(250, 250, 125).
		Style("fill", "none").
		Style("stroke", "black")
	_ = canvas.Write(w)
}
