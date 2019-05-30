package window

import (
	"math"
	"testing"
)

func TestHann(t *testing.T) {
	// SciPy: numpy.set_printoptions(precision=15); scipy.signal.hann(16)
	v := []float64{
		0., 0.0432272711787, 0.165434696820571,
		0.345491502812526, 0.552264231633827, 0.75,
		0.904508497187474, 0.989073800366903, 0.989073800366903,
		0.904508497187474, 0.75, 0.552264231633827,
		0.345491502812526, 0.165434696820571, 0.0432272711787,
		0.,
	}
	n := len(v)

	w, _ := Hann(n)

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (v[i] - w[i]) * (v[i] - w[i])
	}

	mse := math.Sqrt(sum / float64(n))
	if mse > 1e-15 {
		t.Error("window error, MSE:", mse)
	}
}

func TestHamming(t *testing.T) {
	// SciPy: numpy.set_printoptions(precision=15); scipy.signal.hamming(16)
	v := []float64{
		0.08, 0.119769089484404, 0.232199921074925,
		0.397852182587524, 0.588083093103121, 0.77,
		0.912147817412476, 0.989947896337551, 0.989947896337551,
		0.912147817412476, 0.77, 0.588083093103121,
		0.397852182587524, 0.232199921074925, 0.119769089484404,
		0.08,
	}
	n := len(v)
	w, _ := Hamming(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (v[i] - w[i]) * (v[i] - w[i])
	}

	mse := math.Sqrt(sum / float64(n))
	if mse > 1e-15 {
		t.Error("window error, MSE:", mse)
	}
}

func TestBlackman(t *testing.T) {
	// SciPy: numpy.set_printoptions(precision=15); scipy.signal.blackman(16)
	v := []float64{
		-1.387778780781446e-17, 1.675771968740822e-02,
		7.707241975915854e-02, 2.007701432625305e-01,
		3.940124235751222e-01, 6.299999999999999e-01,
		8.492298567374694e-01, 9.821574369783108e-01,
		9.821574369783108e-01, 8.492298567374694e-01,
		6.300000000000002e-01, 3.940124235751227e-01,
		2.007701432625306e-01, 7.707241975915853e-02,
		1.675771968740816e-02, -1.387778780781446e-17,
	}
	n := len(v)
	w, _ := Blackman(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (v[i] - w[i]) * (v[i] - w[i])
	}

	mse := math.Sqrt(sum / float64(n))
	if mse > 1e-15 {
		t.Error("window error, MSE:", mse)
	}
}

func TestBlackmanHarris(t *testing.T) {
	// SciPy: numpy.set_printoptions(precision=15); scipy.signal.blackmanharris(16)
	v := []float64{
		6.000000000000102e-05, 3.600342059774550e-03,
		2.670175342487845e-02, 1.030114893456638e-01,
		2.679881918029911e-01, 5.205749999999999e-01,
		7.938335106543363e-01, 9.748847127123560e-01,
		9.748847127123561e-01, 7.938335106543363e-01,
		5.205750000000002e-01, 2.679881918029914e-01,
		1.030114893456638e-01, 2.670175342487850e-02,
		3.600342059774529e-03, 6.000000000000102e-05,
	}
	n := len(v)
	w, _ := BlackmanHarris(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (v[i] - w[i]) * (v[i] - w[i])
	}

	mse := math.Sqrt(sum / float64(n))
	if mse > 1e-15 {
		t.Error("window error, MSE:", mse)
	}
}

func TestBartlett(t *testing.T) {
	// SciPy: numpy.set_printoptions(precision=15); scipy.signal.bartlett(16)
	v := []float64{
		0., 0.133333333333333, 0.266666666666667,
		0.4, 0.533333333333333, 0.666666666666667,
		0.8, 0.933333333333333, 0.933333333333333,
		0.8, 0.666666666666667, 0.533333333333333,
		0.4, 0.266666666666667, 0.133333333333333,
		0.,
	}
	n := len(v)
	w, _ := Bartlett(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (v[i] - w[i]) * (v[i] - w[i])
	}

	mse := math.Sqrt(sum / float64(n))
	if mse > 1e-15 {
		t.Error("window error, MSE:", mse)
	}
}
