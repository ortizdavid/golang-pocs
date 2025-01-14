package core

import (
	"fmt"
)

type Polygon struct {
	Height float64
	Width  float64
}

func (p *Polygon) hasSameValues(height float64, width float64) bool {
	return p.Height == height && p.Width == width
}


func (p *Polygon) CalculateAny() float64 {
	return  ((p.Height * p.Width) + 0.5 * (p.Height)) 
}

func (p *Polygon) Verify(height float64, width float64) {
	if p.hasSameValues(height, width) {
		fmt.Println("Same Values")
	} else {
		fmt.Println("Different Values")
	}
}

