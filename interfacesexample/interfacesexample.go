// interfacesexample has a short demo on interfaces.
// the wording implements is not used
package interfacesexample

import (
	"log"
	"math"
)

type shape interface {
	perimeter() float64
	// When the line area() is uncommented, b and c no longer implement
	// the shape interface because they don't have area()
	//area() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// equilateral triangle
type eqlTriangle struct {
	side float64
}

func findPerim(s shape) float64 {
	perim := s.perimeter()
	return perim
}

func (r rectangle) perimeter() float64 {
	perim := r.height + r.height + r.width + r.width
	return perim
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c circle) perimeter() float64 {
	perim := 2 * math.Phi * c.radius
	return perim
}

func (t eqlTriangle) perimeter() float64 {
	perim := t.side * 3
	return perim
}

func DemoInterfaces() {
	log.Println("Demo interfaces.")
	a := rectangle{10, 5}
	b := circle{7}
	c := eqlTriangle{3}
	log.Println("a perimeter direct call:", a.perimeter())
	log.Println("b perimeter direct call:", b.perimeter())
	log.Println("c perimeter direct call:", c.perimeter())
	log.Println("interface findperim(a) :", findPerim(a))
	log.Println("interface findperim(b) :", findPerim(b))
	log.Println("interface findperim(c) :", findPerim(c))
	// When the line area() is uncommented, b and c no longer implement
	// the shape interface because they don't have area()
	log.Println("a area direct call:", a.area())
}
