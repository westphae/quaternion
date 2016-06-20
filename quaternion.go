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

type Quaternion struct {
	W, X, Y, Z float64
}

func Conj(qin Quaternion) Quaternion {
	qout := Quaternion{}
	qout.W = +qin.W
	qout.X = -qin.X
	qout.Y = -qin.Y
	qout.Z = -qin.Z
	return qout
}

func Norm2(qin Quaternion) float64 {
	return qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z
}

func Norm(qin Quaternion) float64 {
	return math.Sqrt(qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z)
}

func Scalar(w float64) Quaternion {
	return Quaternion{W: w}
}

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

func Unit(qin Quaternion) Quaternion {
	k := Norm(qin)
	return Quaternion{qin.W / k, qin.X / k, qin.Y / k, qin.Z / k}
}

func Inv(qin Quaternion) Quaternion {
	k2 := Norm2(qin)
	q := Conj(qin)
	return Quaternion{q.W / k2, q.X / k2, q.Y / k2, q.Z / k2}
}
