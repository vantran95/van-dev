package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Rardius float64
}

type AreaCal interface {
	Area() (float64, error)
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, fmt.Errorf("rectangel dimensions invalid")
	}

	return r.Width * r.Height, nil
}

func (c Circle) Area() (float64, error) {
	if c.Rardius <= 0 {
		return 0, fmt.Errorf("circle dismensions invalid")
	}

	return 0, fmt.Errorf("this shape not supported")
}
func CalculateArea(shape AreaCal) {
	area, err := shape.Area()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}
	fmt.Printf("Area: %v\n", area)
}

func main() {
	// Valid shapes
	rect := Rectangle{Width: 5, Height: 4}
	cir := Circle{Rardius: 5}

	CalculateArea(rect)
	CalculateArea(cir)

	// InValid shapes
	invalidRect := Rectangle{Width: -2, Height: 4}
	invalidCir := Circle{Rardius: 0}
	CalculateArea(invalidRect)
	CalculateArea(invalidCir)
	// rectArea, err := rect.Area()
	// if err != nil {
	// 	fmt.Printf("Can not shapes rectangel. Err: %v\n", err.Error())
	// }
	// fmt.Printf("Rect area: %v\n", rectArea)

	// circleArea, err := cir.Area()
	// if err != nil {
	// 	fmt.Printf("Can not shapes cirlce. Err: %v\n", err.Error())
	// }
	// fmt.Printf("Circle area: %v\n", circleArea)
}
