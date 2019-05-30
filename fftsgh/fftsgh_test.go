package fftsgh

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkRdft2(b *testing.B) {
	n := 2
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft4(b *testing.B) {
	n := 4
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft8(b *testing.B) {
	n := 8
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft16(b *testing.B) {
	n := 16
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft32(b *testing.B) {
	n := 32
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft64(b *testing.B) {
	n := 64
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}
func BenchmarkRdft128(b *testing.B) {
	n := 128
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}
func BenchmarkRdft256(b *testing.B) {
	n := 256
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}
func BenchmarkRdft512(b *testing.B) {
	n := 512
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft1024(b *testing.B) {
	n := 1024
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkRdft2048(b *testing.B) {
	n := 2048
	a := make([]float64, n)
	a[0] = 1
	orig := make([]float64, n)
	copy(orig, a)

	for i := 0; i < b.N; i++ {
		Rdft(n, 1, a)
		copy(a, orig)
	}
}

func BenchmarkNumpyRfft512(b *testing.B) {
	n := 512
	a := make([]float64, n)
	a[0] = 1

	for i := 0; i < b.N; i++ {
		NumpyRfft(a)
	}
}

func BenchmarkNumpyRfft1024(b *testing.B) {
	n := 1024
	a := make([]float64, n)
	a[0] = 1

	for i := 0; i < b.N; i++ {
		NumpyRfft(a)
	}
}

func BenchmarkScipyDctNormOrtho32(b *testing.B) {
	signal := make([]float64, 32)
	signal[0] = 1

	for i := 0; i < b.N; i++ {
		ScipyDct(signal, DctNormOrtho)
	}

}

func TestRoundTrip(t *testing.T) {
	for k := 0; k < 16; k++ {
		n := 2 << uint(k)
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = rand.Float64()
		}
		orig := make([]float64, len(a))
		copy(orig, a)
		Rdft(n, 1, a)
		Rdft(n, -1, a)
		sum := 0.0
		for i := 0; i < n; i++ {
			a[i] *= 2.0 / float64(n)
			sum += (orig[i] - a[i]) * (orig[i] - a[i])
		}
		mse := math.Sqrt(sum / float64(n))
		if mse > 1e-15 {
			t.Error("MSE too large for size", mse, n)
		}
	}
}

func TestRdft(t *testing.T) {
	// Test vector
	a := []float64{1, 0, 0, 0, 0, 0, 0, 0}
	// Verification vector
	v := []float64{1, 1, 1, 0, 1, 0, 1, 0}
	n := len(a)
	Rdft(n, 1, a)
	for i := 0; i < n/2; i += 2 {
		if a[i] != v[i] {
			t.Error("Result does not equal test vector")
		}
	}
}

func TestNumpyRfft(t *testing.T) {
	// Test vector
	sig := []float64{1, 0, 0, 0, 0, 0, 0, 0}
	// Verification vector
	// NumPy: numpy.fft.rfft([1, 0, 0, 0, 0, 0, 0, 0])
	ver := []complex128{complex(1, 0), complex(1, 0), complex(1, 0), complex(1, 0), complex(1, 0)}
	fft := NumpyRfft(sig)
	for i := range ver {
		if ver[i] != fft[i] {
			t.Error("Result does not equal test vector")
		}
	}
}
