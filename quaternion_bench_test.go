package quaternion

import "testing"

var (
	benchQ = Quaternion{0.24765262787484427, 0.2940044459739585, 0.3943046179925829, 0.8347175749221727}
	benchR = Quaternion{-0.7904669075670613, 0.44891659738265544, -0.3627631346111533, 0.205033813803568}
	benchV = Vec3{0.3, -0.7, 1.2}

	sinkQ Quaternion
	sinkM [3][3]float64
	sinkV Vec3
	sinkF float64
)

func BenchmarkFromEuler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkQ = FromEuler(-1.2, 0.4, 5.5)
	}
}

func BenchmarkEuler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkF, _, _ = benchQ.Euler()
	}
}

func BenchmarkProd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkQ = Prod(benchQ, benchR)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkQ = benchQ.Mul(benchR)
	}
}

func BenchmarkRotMat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkM = benchQ.RotMat()
	}
}

func BenchmarkRotMatUnit(b *testing.B) {
	q := benchQ.Unit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkM = q.RotMatUnit()
	}
}

func BenchmarkRotateVec3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkV = benchQ.RotateVec3(benchV)
	}
}

func BenchmarkRotateVec3Unit(b *testing.B) {
	q := benchQ.Unit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkV = q.RotateVec3Unit(benchV)
	}
}
