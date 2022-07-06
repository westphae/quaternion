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

// New returns a new quaternion
func New(w, x, y, z float64) Quaternion {
	return Quaternion{W: w, X: x, Y: y, Z: z}
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

// Norm2 returns the L2-Norm of a Quaternion (W,X,Y,Z) -> W*W+X*X+Y*Y+Z*Z
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

// Norm returns the L1-Norm of a Quaternion (W,X,Y,Z) -> Sqrt(W*W+X*X+Y*Y+Z*Z)
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
func (qin Quaternion) Unit() Quaternion {
	k := qin.Norm()
	return Quaternion{qin.W / k, qin.X / k, qin.Y / k, qin.Z / k}
}

// Inv returns the Quaternion conjugate rescaled so that Q Q* = 1
func (qin Quaternion) Inv() Quaternion {
	k2 := qin.Norm2()
	q := qin.Conj()
	return Quaternion{q.W / k2, q.X / k2, q.Y / k2, q.Z / k2}
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
	q := Quaternion{}
	q.W = math.Cos(phi/2)*math.Cos(theta/2)*math.Cos(psi/2) +
		math.Sin(phi/2)*math.Sin(theta/2)*math.Sin(psi/2)
	q.X = math.Sin(phi/2)*math.Cos(theta/2)*math.Cos(psi/2) -
		math.Cos(phi/2)*math.Sin(theta/2)*math.Sin(psi/2)
	q.Y = math.Cos(phi/2)*math.Sin(theta/2)*math.Cos(psi/2) +
		math.Sin(phi/2)*math.Cos(theta/2)*math.Sin(psi/2)
	q.Z = math.Cos(phi/2)*math.Cos(theta/2)*math.Sin(psi/2) -
		math.Sin(phi/2)*math.Sin(theta/2)*math.Cos(psi/2)
	return q
}

// RotMat returns the rotation matrix (as float array) corresponding to a Quaternion
func (qin Quaternion) RotMat() [3][3]float64 {
	q := qin.Unit()
	m := [3][3]float64{}
	m[0][0] = 1 - 2*(q.Y*q.Y+q.Z*q.Z)
	m[0][1] = 2 * (q.X*q.Y - q.W*q.Z)
	m[0][2] = 2 * (q.W*q.Y + q.X*q.Z)

	m[1][1] = 1 - 2*(q.Z*q.Z+q.X*q.X)
	m[1][2] = 2 * (q.Y*q.Z - q.W*q.X)
	m[1][0] = 2 * (q.W*q.Z + q.Y*q.X)

	m[2][2] = 1 - 2*(q.X*q.X+q.Y*q.Y)
	m[2][0] = 2 * (q.Z*q.X - q.W*q.Y)
	m[2][1] = 2 * (q.W*q.X + q.Z*q.Y)
	return m
}

func (a Vec3) Normalize() Vec3 {
	r := 1 / math.Sqrt(float64(a.X*a.X+a.Y*a.Y+a.Z*a.Z))
	return Vec3{a.X * r, a.Y * r, a.Z * r}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Cross(b Vec3) Vec3 {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vec3{x, y, z}
}

func (a Vec3) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Returns a quaternion representing a rotation between the two arbitrary vectors.
// Handles special cases too.
// Refer to:
// https://stackoverflow.com/questions/1171849/finding-quaternion-representing-the-rotation-from-one-vector-to-another#comment20591681_1171995
// Also:
// https://github.com/toji/gl-matrix/blob/f0583ef53e94bc7e78b78c8a24f09ed5e2f7a20c/src/gl-matrix/quat.js#L54
func From2Vecs(v1, v2 Vec3) Quaternion {
	// Required: both vectors are unit length:
	v1 = v1.Normalize()
	v2 = v2.Normalize()

	q := Quaternion{}

	xUnitVec := Vec3{X: 1, Y: 0, Z: 0}
	yUnitVec := Vec3{X: 0, Y: 1, Z: 0}

	dot := v1.Dot(v2)
	if dot < -0.999999 { // Handle the case of parallel vectors pointing in opposite directions.
		tmpvec := xUnitVec.Cross(v1)
		if tmpvec.Length() < 0.000001 {
			tmpvec = yUnitVec.Cross(v1)
		}
		tmpvec = tmpvec.Normalize()
		q.X = float64(tmpvec.X)
		q.Y = float64(tmpvec.Y)
		q.Z = float64(tmpvec.Z)
		q.W = math.Pi
	} else if dot > 0.999999 { // Handle the case of parallel vectors both in the same direction.
		q.X = 0
		q.Y = 0
		q.Z = 0
		q.W = 1
	} else {
		tmpvec := v1.Cross(v2)
		q.X = float64(tmpvec.X)
		q.Y = float64(tmpvec.Y)
		q.Z = float64(tmpvec.Z)
		// Note: this statement is 1, if both vectors are unit length:
		// 1 == math.Sqrt(float64(v1.LengthSquared())*float64(v2.LengthSquared()))
		q.W = 1 + float64(dot)
	}

	// Don't forget to normalize q.
	q = q.Unit()

	return q
}
