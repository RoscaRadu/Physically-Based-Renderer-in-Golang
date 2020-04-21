package trace

import (
	"database/sql"
	"math"
	"math/rand"
)

type Vector3 struct{
	X,Y,Z float64
}
//Vector scaling by n
func (a Vector3) Scale (n float64) Vector3{
	return Vector3{a.X*n,a.Y*n,a.Z*n}
}

//Multiplying two vectors
func (a Vector3) Mult(b Vector3) Vector3{
	return Vector3{a.X*b.X,a.Y*b.Y,a.Z*b.Z}
}

//Adding two vectors
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

//
