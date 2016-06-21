/*
Golang package implementing quaternion math
Purpose is to provide quaternion support under the MIT license as existing
Go quaternion packages are under more restrictive or unspecified licenses.

This project is licensed under the terms of the MIT license.
*/

package quaternion

import (
	"math"
)

// Quaternion represents a quaternion W+X*i+Y*j+Z*k
type Quaternion struct {
	W float64 // Scalar component
	X float64 // i component
	Y float64 // j component
	Z float64 // k component
}

// Conj returns the conjugate of a Quaternion (W,X,Y,Z) -> (W,-X,-Y,-Z)
func Conj(qin Quaternion) Quaternion {
	qout := Quaternion{}
	qout.W = +qin.W
	qout.X = -qin.X
	qout.Y = -qin.Y
	qout.Z = -qin.Z
	return qout
}

// Norm2 returns the L2-Norm of a Quaternion (W,X,Y,Z) -> W*W+X*X+Y*Y+Z*Z
func Norm2(qin Quaternion) float64 {
	return qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z
}

// Norm returns the L1-Norm of a Quaternion (W,X,Y,Z) -> Sqrt(W*W+X*X+Y*Y+Z*Z)
func Norm(qin Quaternion) float64 {
	return math.Sqrt(qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z)
}

// Scalar returns a scalar-only Quaternion representation of a float (W,0,0,0)
func Scalar(w float64) Quaternion {
	return Quaternion{W: w}
}

// Sum returns the vector sum of any number of Quaternions
func Sum(qin ...Quaternion) Quaternion {
	qout := Quaternion{}
	for _, q := range qin {
		qout.W += q.W
		qout.X += q.X
		qout.Y += q.Y
		qout.Z += q.Z
	}
	return qout
}

// Prod returns the non-commutative product of any number of Quaternions
func Prod(qin ...Quaternion) Quaternion {
	qout := Quaternion{1, 0, 0, 0}
	var w, x, y, z float64
	for _, q := range qin {
		w = qout.W*q.W - qout.X*q.X - qout.Y*q.Y - qout.Z*q.Z
		x = qout.W*q.X + qout.X*q.W + qout.Y*q.Z - qout.Z*q.Y
		y = qout.W*q.Y + qout.Y*q.W + qout.Z*q.X - qout.X*q.Z
		z = qout.W*q.Z + qout.Z*q.W + qout.X*q.Y - qout.Y*q.X
		qout = Quaternion{w, x, y, z}
	}
	return qout
}

// Unit returns the Quaternion rescaled to unit-L1-norm
func Unit(qin Quaternion) Quaternion {
	k := Norm(qin)
	return Quaternion{qin.W / k, qin.X / k, qin.Y / k, qin.Z / k}
}

// Inv returns the Quaternion conjugate rescaled so that Q Q* = 1
func Inv(qin Quaternion) Quaternion {
	k2 := Norm2(qin)
	q := Conj(qin)
	return Quaternion{q.W / k2, q.X / k2, q.Y / k2, q.Z / k2}
}
