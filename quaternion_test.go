/*
Golang package implementing quaternion math
Purpose is to provide quaternion support under the MIT license as existing
Go quaternion packages are under more restrictive or unspecified licenses.

This project is licensed under the terms of the MIT license.
*/

package quaternion

import (
	"math"
	"testing"
)

var (
	qs1 = Quaternion{1, 0, 0, 0}
	qs2 = Quaternion{W: 10}
	qs3 = Scalar(11)
	qs4 = Scalar(110)
	qv1 = Quaternion{0, 1, 0, 0}
	qv2 = Quaternion{0, 0, 1, 1}
	qv3 = Quaternion{0, 1, 0, 1}
	qv4 = Quaternion{0, 2, 1, 2}
	qv5 = Quaternion{0, -2, -1, -2}
	qv6 = Quaternion{-1, -1, 1, 1}
	q1  = Quaternion{1, -1, -1, 1}
	q2  = Quaternion{-1, 1, 1, -1}
	q0  = Quaternion{0, 0, 0, 0}
	q3  = Quaternion{-1, -1, -1, 1}
	q4  = Quaternion{4, -4, -4, 4}
	q5  = Quaternion{0.5, -0.5, -0.5, 0.5}
	q6  = Quaternion{0.0625, 0.0625, 0.0625, -0.0625}
	q7  = Quaternion{0.24765262787484427, 0.2940044459739585, 0.3943046179925829, 0.8347175749221727}
	q8  = Quaternion{-0.7904669075670613, 0.44891659738265544, -0.3627631346111533, 0.205033813803568}
	q9  = Quaternion{math.Cos(math.Pi / 2), math.Sin(math.Pi/2) / math.Sqrt(3),
		math.Sin(math.Pi/2) / math.Sqrt(3), -math.Sin(math.Pi/2) / math.Sqrt(3)}
	q10 = Quaternion{0.707106781186548, 0.707106781186547, 0, 0}
	m   = [3][3]float64{[3]float64{-0.333333333, 0.666666667, -0.666666667},
		[3]float64{0.666666667, -0.333333333, -0.666666667},
		[3]float64{-0.666666667, -0.666666667, -0.333333333}}
	v1 = Vec3{0, 0, 1}
	v2 = Vec3{0, -1, 0}
)

func TestScalarSum(t *testing.T) {
	if Sum(qs1, qs2) != qs3 {
		t.Fail()
	}
}

func TestVectorSum(t *testing.T) {
	if Sum(qv1, qv2, qv3) != qv4 {
		t.Fail()
	}
}

func TestMixedSum(t *testing.T) {
	if Sum(q1, q2) != q0 {
		t.Fail()
	}
}

func TestScalarConj(t *testing.T) {
	if qs1.Conj() != qs1 {
		t.Fail()
	}
}

func TestVectorConj(t *testing.T) {
	if qv4.Conj() != qv5 {
		t.Fail()
	}
}

func TestMixedConj(t *testing.T) {
	if q2.Conj() != q3 {
		t.Fail()
	}
}

func TestScalarProd(t *testing.T) {
	if Prod(qs1, qs2, qs3) != qs4 {
		t.Fail()
	}
}

func TestVectorProd(t *testing.T) {
	if Prod(qv1, qv2, qv3) != qv6 {
		t.Fail()
	}
}

func TestMixedProd(t *testing.T) {
	if Prod(q1, q2, q3) != q4 {
		t.Fail()
	}
}

func TestScalarNorm(t *testing.T) {
	if qs4.Norm() != 110 {
		t.Fail()
	}
}

func TestVectorNorm(t *testing.T) {
	if qv4.Norm() != 3 {
		t.Fail()
	}
}

func TestMixedNorm(t *testing.T) {
	if q4.Norm() != 8 {
		t.Fail()
	}
}

func TestUnit(t *testing.T) {
	if q4.Unit() != q5 {
		t.Fail()
	}
}

func TestInv(t *testing.T) {
	if q4.Inv() != q6 {
		t.Fail()
	}
}

func TestRotateVec3(t *testing.T) {
	rot := q10.RotateVec3(v1)
	if math.Abs(rot.X) > 1e-6 ||
		math.Abs(rot.Y+1) > 1e-6 ||
		math.Abs(rot.Z) > 1e-6 {
		t.Fail()
	}
}

func TestVec3Rotate(t *testing.T) {
	vec := v1.Rotate(q10)
	if math.Abs(vec.X) > 1e-6 ||
		math.Abs(vec.Y+1) > 1e-6 ||
		math.Abs(vec.Z) > 1e-6 {
		t.Fail()
	}
}

func TestEuler(t *testing.T) {
	phi, theta, psi := q7.Euler()
	if math.Abs(phi-1.0) > 1e-6 ||
		math.Abs(theta+0.3) > 1e-6 ||
		math.Abs(psi-2.4) > 1e-6 {
		t.Fail()
	}
}

func TestFromEuler(t *testing.T) {
	q := FromEuler(-1.2, 0.4, 5.5)
	if math.Abs(q.W-q8.W) > 1e-6 ||
		math.Abs(q.X-q8.X) > 1e-6 ||
		math.Abs(q.Y-q8.Y) > 1e-6 ||
		math.Abs(q.Z-q8.Z) > 1e-6 {
		t.Fail()
	}

}

func TestRotMat(t *testing.T) {
	mm := q9.RotMat()
	for i, x := range mm {
		for j, y := range x {
			if math.Abs(m[i][j]-y) > 1e-6 {
				t.Fail()
			}
		}
	}
}

func TestFrom2Vecs(t *testing.T) {
	q := From2Vecs(v1, v2)
	t.Logf("Quaternion W: %v, X: %v, Y: %v, Z: %v", q.W, q.X, q.Y, q.Z)
	if math.Abs(q.W-q10.W) > 1e-6 ||
		math.Abs(q.X-q10.X) > 1e-6 ||
		math.Abs(q.Y-q10.Y) > 1e-6 ||
		math.Abs(q.Z-q10.Z) > 1e-6 {
		t.Fail()
	}
}
