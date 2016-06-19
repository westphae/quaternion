# quaternion
Quaternion math in golang

Instantiate a quaternion q = a_w + a_x * **i** + a_y * **j** + a_z * **k**:
```go
q1 := Quaternion{a_w, a_x, a_y, a_z}
q2 := Quaternion{W: 0.5, X: 0.5, Y: -0.707, Z: -0.707}
```

Calculate the sum as a new quaternion:
```go
q3 := Sum(q1, q2)
```

Sum takes any number of quaternions as arguments:
```go
q4 := Sum(q3, q1, q2, q4)
```
