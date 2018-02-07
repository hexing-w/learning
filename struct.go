package fff

import (
	"fmt"
	"math"
)

type Triple struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Triple) area() float64 {
	return r.width * r.height / 2
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func main() {
	r1 := Triple{3, 6}
	r2 := Circle{4}
	fmt.Println("area is %d", r1.area())
	fmt.Println("area is %s", r2.area())
}
