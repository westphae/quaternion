/*
Golang package implementing quaternion math
Purpose is to provide quaternion support under the MIT license as existing
Go quaternion packages are under more restrictive or unspecified licenses.

This project is licensed under the terms of the MIT license.
*/

package quaternion

import (
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
	if Conj(qs1) != qs1 {
		t.Fail()
	}
}

func TestVectorConj(t *testing.T) {
	if Conj(qv4) != qv5 {
		t.Fail()
	}
}

func TestMixedConj(t *testing.T) {
	if Conj(q2) != q3 {
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
	if Norm(qs4) != 110 {
		t.Fail()
	}
}

func TestVectorNorm(t *testing.T) {
	if Norm(qv4) != 3 {
		t.Fail()
	}
}

func TestMixedNorm(t *testing.T) {
	if Norm(q4) != 8 {
		t.Fail()
	}
}
