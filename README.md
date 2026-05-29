# quaternion
Quaternion math in golang

Note: There are other packages to support quaternions in Go.  See, for example,
https://github.com/thisisneal/Quaternion.  I couldn't find any under an
unrestrictive license, so this is an implementation under the MIT license.

## Quaternions
Quaternions are defined by **i** * **i** = **j** * **j** = **k** * **k** = **i** * **j** * **k** = -1.
See https://en.wikipedia.org/wiki/Quaternion

## Usage
Instantiate a quaternion q = a_w + a_x * **i** + a_y * **j** + a_z * **k**:
```Go
q1 := Quaternion{a_w, a_x, a_y, a_z}
q2 := Quaternion{W: 0.5, X: 0.5, Y: -0.707, Z: -0.707}
q3 := New(0.5,0.5,-0.707,-0.707)
```

A number (scalar) can be created with Quaternion or with the Scalar function:
```Go
qk1 := Quaternion{W: -0.5}
qk2 := Scalar(0.5)
```

A pure quaternion with no scalar component:
```Go
q := Pure(0.5, -0.707, -0.707)
```

The identity ("no rotation") quaternion:
```Go
q := Identity()
```

Calculate the conjugate q* = a_w - a_x * **i** - a_y * **j** - a_z * **k**:
```Go
q5 := qr.Conj()
```

Calculate the sum as a new quaternion:
```Go
q3 := Sum(q1, q2)
```

Sum takes any number of quaternions as arguments:
```Go
q4 := Sum(q3, q1, q2, q4)
```

Prod works the same way as Sum:
```Go
q5 := Prod(q4, q3, q1, q2)
```

For the common two-quaternion case, Mul is a faster binary product.
Scale multiplies by a scalar and Sub takes a difference:
```Go
q6 := q1.Mul(q2)
q7 := q1.Scale(0.5)
q8 := q1.Sub(q2)
d  := q1.Dot(q2)
```

Calculate the norm ("length") and the squared norm:
```Go
k := q5.Norm()
k2 := q5.Norm2()
```

Convert to/from Euler angle representations:
```Go
q1 := FromEuler(math.Pi/4, math.Pi/3, 5*math.Pi/3)
phi, theta, psi := q1.Euler()
```

Convert to/from axis-angle representations:
```Go
q := FromAxisAngle(Vec3{0, 0, 1}, math.Pi/2)
axis, angle := q.AxisAngle()
```

Interpolate between two rotations. Slerp follows the shortest arc at constant
angular velocity; Nlerp is a cheaper approximation:
```Go
q := Slerp(q1, q2, 0.5)
q := Nlerp(q1, q2, 0.5)
```

Rotate a vector by a quaternion, and get the Rotation Matrix:
```Go
v := q1.RotateVec3(Vec3{0, 0, 1})
m := q1.RotMat()
```

`RotateVec3` and `RotMat` normalize the quaternion first. If you already hold a
unit quaternion, the `Unit` variants skip that step and run roughly twice as
fast:
```Go
v := q1.RotateVec3Unit(Vec3{0, 0, 1})
m := q1.RotMatUnit()
```
