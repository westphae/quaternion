/*
Golang package implementing quaternion math
Purpose is to provide quaternion support under the MIT license as existing
Go quaternion packages are under more restrictive or unspecified licenses.

This project is licensed under the terms of the MIT license.
*/

package quaternion

import ()

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
