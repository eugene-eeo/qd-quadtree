package main

import "math"

type Vector struct {
	I, J float64
}

func NewVectorFromPoints(A, B *Point) *Vector {
	return &Vector{
		I: A.X - B.X,
		J: A.Y - B.Y,
	}
}

func (a *Vector) Magnitude() float64 {
	return math.Sqrt(a.I*a.I + a.J*a.J)
}

func (a *Vector) Dot(b *Vector) float64 {
	return a.I*b.I + a.J*b.J
}

func (a *Vector) AngleBetween(b *Vector) float64 {
	return math.Acos(a.Dot(b) / (a.Magnitude() * b.Magnitude()))
}
