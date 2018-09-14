package syntax

import (
	"math"
	"fmt"
)

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type Triangle struct {

}

func (ci *Triangle) Area() float32 {
	return 0
}

func TypeAssert() {
	var area Shape
	sq1 := new(Square)
	sq1.side = 5

	area = new(Triangle)
	// Is Square the type of area?
	if t, ok := area.(*Square); ok {
		fmt.Printf("The type of area is: %T\n", t)
	}
	if u, ok := area.(*Circle); ok {
		fmt.Printf("The type of area is: %T\n", u)
	} else {
		fmt.Println("area does not contain a variable of type Circle")
	}
}
