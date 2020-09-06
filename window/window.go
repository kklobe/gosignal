package window

import (
	"errors"
	"math"
)

// Hann returns an n-sized Hann window
func Hann(n int) ([]float64, error) {
	if n < 1 {
		return nil, errors.New("invalid window size")
	}
	w := make([]float64, n)
	if n == 1 {
		w[0] = 1
	} else {
		k := 2 * math.Pi / float64(n-1)

		for i := range w {
			w[i] = 0.5 * (1 - math.Cos(k*float64(i)))
		}
	}
	return w, nil
}

// Hamming returns an n-sized Hamming window
func Hamming(n int) ([]float64, error) {
	if n < 1 {
		return nil, errors.New("invalid window size")
	}
	w := make([]float64, n)
	if n == 1 {
		w[0] = 1
	} else {
		k := 2 * math.Pi / float64(n-1)

		for i := range w {
			w[i] = 0.54 -
				0.46*math.Cos(k*float64(i))
		}
	}
	return w, nil
}

// Blackman returns an n-sized Blackman window
func Blackman(n int) ([]float64, error) {
	if n < 1 {
		return nil, errors.New("invalid window size")
	}
	w := make([]float64, n)
	for i := range w {
		w[i] = 0.42 -
			0.50*math.Cos((2*math.Pi*float64(i))/float64(n-1)) +
			0.08*math.Cos((4*math.Pi*float64(i))/float64(n-1))
	}
	return w, nil
}

// BlackmanHarris returns an n-sized Blackman-Harris window
func BlackmanHarris(n int) ([]float64, error) {
	if n < 1 {
		return nil, errors.New("invalid window size")
	}
	w := make([]float64, n)
	if n == 1 {
		w[0] = 1
	} else {
		k := 2 * math.Pi / float64(n-1)

		for i := range w {
			w[i] = 0.35875 -
				0.48829*math.Cos(k*float64(i)) +
				0.14128*math.Cos(2*k*float64(i)) -
				0.01168*math.Cos(3*k*float64(i))
		}
	}
	return w, nil
}

// Bartlett returns an n-sized Bartlett window
func Bartlett(n int) ([]float64, error) {
	if n < 1 {
		return nil, errors.New("invalid window size")
	}
	w := make([]float64, n)
	if n == 1 {
		w[0] = 1
	} else {
		n0 := float64(n - 1)
		for i := range w {
			i0 := float64(i)
			w[i] = 2 / n0 * (n0/2 - math.Abs(i0-n0/2))
		}
	}
	return w, nil
}
