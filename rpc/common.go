package rpc

import (
	"math"
)

type Point struct {
	X, Y int
}

type Line struct {
	A, B Point
}

type Rect struct {
	Width, Height int
}

type Circle struct {
	P Point
	R int
}

func (t *Point) Quadrant(p Point, quad *string) error {
	if p.X < 0 {
		if p.Y < 0 {
			*quad = "Third Quadrant"
		} else if p.Y == 0 {
			*quad = "X Direct Line"
		} else {
			*quad = "Second Quadrant"
		}
	} else if p.X > 0 {
		if p.Y < 0 {
			*quad = "Fourth Quadrant"
		} else if p.Y == 0 {
			*quad = "X Direct Line"
		} else {
			*quad = "First Quadrant"
		}
	}
	return nil
}

/*
func (t *Line) IsOnLine(l Line, p Point, on *bool) error {
	*on = (l.A.X-p.X)*(p.Y-l.B.Y) == (l.A.Y-p.Y)*(p.X-l.B.X)
	return nil
}*/

func (t *Line) Length(l Line, leng *float64) error {
	*leng = math.Sqrt(float64((l.A.X-l.B.X)*(l.A.X-l.B.X)) + float64((l.A.Y-l.B.Y)*(l.A.Y-l.B.Y)))
	return nil
}

func (r *Rect) Area(rect Rect, area *float64) error {
	*area = float64(rect.Width * rect.Height)
	return nil
}

func (c *Circle) Area(circle Circle, area *float64) error {
	*area = math.Pi * float64(circle.R*circle.R)
	return nil
}
