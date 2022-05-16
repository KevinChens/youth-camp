package _interface

import "testing"

func TestMeasure(t *testing.T) {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}
	Measure(r)
	Measure(c)
}