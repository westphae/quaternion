# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

A single-package Go library (`package quaternion`, module `github.com/westphae/quaternion`) implementing quaternion math under the MIT license. It exists specifically because other Go quaternion packages were found only under more restrictive or unspecified licenses. All code lives in two files: `quaternion.go` (implementation) and `quaternion_test.go` (tests).

## Commands

```bash
go build ./...                    # build
go test ./...                     # run all tests
go test -run TestRotMat           # run a single test by name
go test -v ./...                  # verbose test output
go vet ./...                      # vet
gofmt -w quaternion.go            # format
```

## Architecture & conventions

- **Two core types:** `Quaternion{W, X, Y, Z}` (W is the scalar component; X/Y/Z are the i/j/k components) and `Vec3{X, Y, Z}`. Both are plain value structs compared directly with `==` in tests.

- **Value semantics throughout.** Methods take a value receiver, mutate the local copy, and return it (e.g. `Conj`, `Neg`). Nothing mutates in place; every operation returns a new value. Preserve this style — don't introduce pointer receivers or in-place mutation.

- **Products.** `Prod(...)` (variadic) and `Mul(r)` (binary) implement the non-commutative Hamilton product — order matters. `Prod` folds left from the identity `{1,0,0,0}` by calling `Mul`, so the product math lives in one place (`Mul`); prefer `Mul` for the common two-quaternion case. `Sum(...)` is the variadic component-wise sum; `Sub`/`Scale`/`Dot` round out the algebra. Constructors: `New`, `Pure` (zero scalar), `Scalar` (zero vector), `Identity()`.

- **Rotation is done two equivalent ways:** `Quaternion.RotateVec3(Vec3)` and `Vec3.Rotate(Quaternion)` — the latter delegates to the former, so there's a single implementation. It uses the optimized cross-product formula on a normalized quaternion (`v + 2w(qv×v) + 2(qv×(qv×v))`), so the result is always a pure rotation regardless of input norm.

- **`…Unit` fast paths:** `RotateVec3Unit`/`RotateUnit` and `RotMatUnit` skip the internal `.Unit()` normalization and assume a unit quaternion (~2× faster; meaningful only for unit input). `RotateVec3` delegates to `RotateVec3Unit` (single impl), but `RotMat` and `RotMatUnit` intentionally duplicate their matrix-build body — delegating made the safe path copy the 72-byte array twice and regressed it; `TestRotMatUnit` asserts the two stay equal, so duplication is guarded. Fast variants exist only where normalization is a meaningful fraction of cost — `Euler`/`AxisAngle`/`Slerp` are dominated by `Atan2`/`Asin`/`Acos`, so none was added.

- **Normalization:** `Euler()`, `RotMat()`, `AxisAngle()`, and the rotation/interpolation funcs operate on the unit form (they call `.Unit()` internally), so they only make sense for rotation quaternions. `Unit()` and `Inv()` guard the zero quaternion (returned unchanged) to avoid `Inf`/`NaN`. `Norm`/`Norm2` are the Euclidean norm and its square.

- **Rotation representations & interpolation:** quaternion ↔ Euler (`FromEuler`/`Euler`), ↔ axis-angle (`FromAxisAngle`/`AxisAngle`), → rotation matrix (`RotMat`). Interpolate with `Slerp` (shortest-arc, constant angular velocity, falls back to `Nlerp` near-parallel) or the cheaper `Nlerp`. Both take the shortest path by negating an input when the dot product is negative.

- **`Vec3` is a full small vector type:** `Add`, `Sub`, `Scale`, `Dot`, `Cross`, `Norm`, `Normalize` (plus `Rotate`). These support axis-angle and the rotation formula; reuse them rather than inlining component arithmetic.

- **Euler convention:** `FromEuler(phi, theta, psi)` / `Euler()` use the aerospace roll-pitch-yaw (Tait-Bryan) formulas. If you touch these, verify against `TestEuler`/`TestFromEuler`, which pin exact expected values.

## Testing notes

Tests use hand-computed expected quaternions. Exact-equality tests (`!=`) are used for operations that are exact in floating point (sums, conjugates, integer-valued products/norms); tolerance comparisons (`math.Abs(...) > 1e-6`) are used for trig-based operations (rotation, Euler, rotation matrix). Follow this split when adding tests — use `1e-6` tolerance for anything involving `math.Sin`/`Cos`/`Sqrt`/`Atan2`.
