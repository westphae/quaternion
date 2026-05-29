/*
Golang package implementing quaternion math
Purpose is to provide quaternion support under the MIT license as existing
Go quaternion packages are under more restrictive or unspecified licenses.

This project is licensed under the terms of the MIT license.
*/

package quaternion

import (
	"fmt"
	"math"
)

// New returns a new quaternion
func New(w, x, y, z float64) Quaternion {
	return Quaternion{W: w, X: x, Y: y, Z: z}
}

// Identity returns the multiplicative identity quaternion (1,0,0,0),
// which represents "no rotation".
func Identity() Quaternion {
	return Quaternion{W: 1}
}

// Pure returns a new pure quaternion (no scalar part)
func Pure(x, y, z float64) Quaternion {
	return Quaternion{X: x, Y: y, Z: z}
}

// Vec3 represents a vector in 3d space
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Quaternion represents a quaternion W+X*i+Y*j+Z*k
type Quaternion struct {
	W float64 // Scalar component
	X float64 // i component
	Y float64 // j component
	Z float64 // k component
}

// Conj returns the conjugate of a Quaternion (W,X,Y,Z) -> (W,-X,-Y,-Z)
func (qin Quaternion) Conj() Quaternion {
	qin.X = -qin.X
	qin.Y = -qin.Y
	qin.Z = -qin.Z
	return qin
}

// Norm2 returns the squared Euclidean norm of a Quaternion (W,X,Y,Z) -> W*W+X*X+Y*Y+Z*Z
func (qin Quaternion) Norm2() float64 {
	return qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z
}

// Neg returns the negative
func (qin Quaternion) Neg() Quaternion {
	qin.W = -qin.W
	qin.X = -qin.X
	qin.Y = -qin.Y
	qin.Z = -qin.Z
	return qin
}

// Norm returns the Euclidean (L2) norm of a Quaternion (W,X,Y,Z) -> Sqrt(W*W+X*X+Y*Y+Z*Z)
func (qin Quaternion) Norm() float64 {
	return math.Sqrt(qin.Norm2())
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

// Mul returns the non-commutative Hamilton product q*r of two Quaternions
func (q Quaternion) Mul(r Quaternion) Quaternion {
	return Quaternion{
		W: q.W*r.W - q.X*r.X - q.Y*r.Y - q.Z*r.Z,
		X: q.W*r.X + q.X*r.W + q.Y*r.Z - q.Z*r.Y,
		Y: q.W*r.Y + q.Y*r.W + q.Z*r.X - q.X*r.Z,
		Z: q.W*r.Z + q.Z*r.W + q.X*r.Y - q.Y*r.X,
	}
}

// Prod returns the non-commutative product of any number of Quaternions
func Prod(qin ...Quaternion) Quaternion {
	if len(qin) == 0 {
		return Identity()
	}
	qout := qin[0]
	for _, q := range qin[1:] {
		qout = qout.Mul(q)
	}
	return qout
}

// Scale returns the Quaternion with every component multiplied by k
func (qin Quaternion) Scale(k float64) Quaternion {
	return Quaternion{qin.W * k, qin.X * k, qin.Y * k, qin.Z * k}
}

// Sub returns the component-wise difference q-r of two Quaternions
func (q Quaternion) Sub(r Quaternion) Quaternion {
	return Quaternion{q.W - r.W, q.X - r.X, q.Y - r.Y, q.Z - r.Z}
}

// Dot returns the dot product of two Quaternions treated as 4-vectors
func (q Quaternion) Dot(r Quaternion) float64 {
	return q.W*r.W + q.X*r.X + q.Y*r.Y + q.Z*r.Z
}

// ApproxEqual reports whether every component of q and r differs by no more than tol
func (q Quaternion) ApproxEqual(r Quaternion, tol float64) bool {
	return math.Abs(q.W-r.W) <= tol &&
		math.Abs(q.X-r.X) <= tol &&
		math.Abs(q.Y-r.Y) <= tol &&
		math.Abs(q.Z-r.Z) <= tol
}

// String implements fmt.Stringer, e.g. "0.5 + 0.5i - 0.707j - 0.707k"
func (q Quaternion) String() string {
	return fmt.Sprintf("%g %+gi %+gj %+gk", q.W, q.X, q.Y, q.Z)
}

// Unit returns the Quaternion rescaled to unit norm. The zero quaternion
// has no unit form and is returned unchanged.
func (qin Quaternion) Unit() Quaternion {
	k := qin.Norm()
	if k == 0 {
		return qin
	}
	return Quaternion{qin.W / k, qin.X / k, qin.Y / k, qin.Z / k}
}

// Inv returns the Quaternion conjugate rescaled so that Q Q* = 1. The zero
// quaternion has no inverse and is returned unchanged.
func (qin Quaternion) Inv() Quaternion {
	k2 := qin.Norm2()
	if k2 == 0 {
		return qin
	}
	q := qin.Conj()
	return Quaternion{q.W / k2, q.X / k2, q.Y / k2, q.Z / k2}
}

// Add returns the vector sum v+w
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub returns the vector difference v-w
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Scale returns the vector with every component multiplied by k
func (v Vec3) Scale(k float64) Vec3 {
	return Vec3{v.X * k, v.Y * k, v.Z * k}
}

// Dot returns the dot product v·w
func (v Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Cross returns the cross product v×w
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// Norm returns the Euclidean length of the vector
func (v Vec3) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}

// Normalize returns the vector rescaled to unit length. The zero vector is
// returned unchanged.
func (v Vec3) Normalize() Vec3 {
	k := v.Norm()
	if k == 0 {
		return v
	}
	return v.Scale(1 / k)
}

// RotateVec3 returns the vector rotated by the quaternion.
func (qin Quaternion) RotateVec3(vec Vec3) Vec3 {
	conj := qin.Conj()
	aug := Quaternion{0, vec.X, vec.Y, vec.Z}
	rot := Prod(qin, aug, conj)
	return Vec3{rot.X, rot.Y, rot.Z}
}

// Rotate returns the vector rotated by the quaternion.
func (vin Vec3) Rotate(q Quaternion) Vec3 {
	conj := q.Conj()
	aug := Quaternion{0, vin.X, vin.Y, vin.Z}
	rot := Prod(q, aug, conj)
	return Vec3{rot.X, rot.Y, rot.Z}
}

// Euler returns the Euler angles phi, theta, psi corresponding to a Quaternion
func (q Quaternion) Euler() (float64, float64, float64) {
	r := q.Unit()
	phi := math.Atan2(2*(r.W*r.X+r.Y*r.Z), 1-2*(r.X*r.X+r.Y*r.Y))
	theta := math.Asin(2 * (r.W*r.Y - r.Z*r.X))
	psi := math.Atan2(2*(r.X*r.Y+r.W*r.Z), 1-2*(r.Y*r.Y+r.Z*r.Z))
	return phi, theta, psi
}

// FromEuler returns a Quaternion corresponding to Euler angles phi, theta, psi
func FromEuler(phi, theta, psi float64) Quaternion {
	// Each half-angle sin/cos is computed once; trig calls dominate the cost.
	cp, sp := math.Cos(phi/2), math.Sin(phi/2)
	ct, st := math.Cos(theta/2), math.Sin(theta/2)
	cs, ss := math.Cos(psi/2), math.Sin(psi/2)
	return Quaternion{
		W: cp*ct*cs + sp*st*ss,
		X: sp*ct*cs - cp*st*ss,
		Y: cp*st*cs + sp*ct*ss,
		Z: cp*ct*ss - sp*st*cs,
	}
}

// FromAxisAngle returns a Quaternion representing a rotation of angle radians
// about the given axis. The axis is normalized internally.
func FromAxisAngle(axis Vec3, angle float64) Quaternion {
	a := axis.Normalize()
	s := math.Sin(angle / 2)
	return Quaternion{
		W: math.Cos(angle / 2),
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

// AxisAngle returns the unit axis and angle (in radians) of the rotation
// represented by the Quaternion. For a rotation of ~0 the axis is undefined
// and (1,0,0) is returned.
func (qin Quaternion) AxisAngle() (axis Vec3, angle float64) {
	q := qin.Unit()
	angle = 2 * math.Acos(q.W)
	s := math.Sqrt(1 - q.W*q.W)
	if s < 1e-9 {
		return Vec3{1, 0, 0}, angle
	}
	return Vec3{q.X / s, q.Y / s, q.Z / s}, angle
}

// RotMat returns the rotation matrix (as float array) corresponding to a Quaternion
func (qin Quaternion) RotMat() [3][3]float64 {
	q := qin.Unit()
	xx, yy, zz := q.X*q.X, q.Y*q.Y, q.Z*q.Z
	xy, xz, yz := q.X*q.Y, q.X*q.Z, q.Y*q.Z
	wx, wy, wz := q.W*q.X, q.W*q.Y, q.W*q.Z
	return [3][3]float64{
		{1 - 2*(yy+zz), 2 * (xy - wz), 2 * (wy + xz)},
		{2 * (wz + xy), 1 - 2*(zz+xx), 2 * (yz - wx)},
		{2 * (xz - wy), 2 * (wx + yz), 1 - 2*(xx+yy)},
	}
}
