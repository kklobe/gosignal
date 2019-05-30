package fftsgh

import "math"

const ( // iota is reset to 0
	DctNormNone  = iota // c0 == 0
	DctNormOrtho = iota // c1 == 1
)

// NumpyRfft matches numpy.fft.rfft behavior
func NumpyRfft(signal []float64) []complex128 {
	n := nextPowerOf2(len(signal))
	s := make([]float64, n)
	copy(s, signal)
	Rdft(n, 1, s)

	c := make([]complex128, len(s)/2+1)
	c[0] = complex(s[0], 0)
	c[len(c)-1] = complex(s[1], 0)
	for i := 1; i < len(c)-1; i++ {
		c[i] = complex(s[2*i], -s[2*i+1])
	}

	return c
}

func nextPowerOf2(n int) int {
	k := 1
	for k < n {
		k *= 2
	}
	return k
}

// ScipyDct matches scipy.fftpack.dct type='2' output
func ScipyDct(x []float64, norm int) []float64 {
	N := len(x)
	y := make([]float64, N)

	for k := 0; k < N; k++ {
		sum := 0.0
		for n := 0; n < N; n++ {
			sum += x[n] * math.Cos(math.Pi*float64(k)*float64(2*n+1)/float64(2*N))
		}
		y[k] = 2.0 * sum
		if norm == DctNormOrtho {
			var f float64
			if k == 0 {
				f = math.Sqrt(1.0 / float64(4*N))
			} else {
				f = math.Sqrt(1.0 / float64(2*N))
			}
			y[k] *= f
		}
	}

	return y
}
