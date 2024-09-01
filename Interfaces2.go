package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type measurable interface {
	perimiter() float64
}

type geometric interface {
	shape
	measurable
}

type rectangle struct {
	width, height float64
}

type circ struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimiter() float64 {
	return 2 * (r.height + r.width)
}

func (c circ) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circ) perimiter() float64 {
	return 2 * math.Pi * c.radius
}

func getGeometry(g geometric) {
	fmt.Println("shape:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimiter())
}

func main() {
	r1 := rectangle{
		height: 10,
		width:  12,
	}
	getGeometry(r1)

	c1 := circ{
		radius: 2,
	}
	getGeometry(c1)

	getSqRt(20)
}

type CalculationErr struct {
	msg string
}

func (ce CalculationErr) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationErr{
			msg: "Invalid input",
		}
	}
	return math.Sqrt(val), nil
}

func getSqRt(val float64) {
	res, err := performCalculation(val)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	fmt.Printf("SQRT of %.2f is %.2f\n", val, res)
}
