# fftsg

fftsg provides a pure Go port of [Ooura FFT's](http://www.kurims.kyoto-u.ac.jp/~ooura/fft.html) fftsg.c. The API is almost identical to the Ooura C version for ease of porting existing code that consumes it, which was the primary motivation.

Some reasons:

- Pure Go for easy cross-compiling
- Unrestricted license
- Good performance

# Notes

I've only created tests and benchmarks for basic Real FFT usage, since that's 100% of my usage. However, I've no reason to believe the other functions won't be identical to the C behavior.

There are extensive comments in the fftsg.go file from the original detailing usage, especially the ip and w parameters.

## Installation and Usage

```$ go get -u github.com/kklobe/gosignal/fftsg```

```
package main

import (
	"fmt"
	"math"
	"github.com/kklobe/gosignal/fftsg"
)

func main() {
	a := []float64{1, 0, 0, 0, 0, 0, 0, 0}
	n := len(a)
	ipLen := int(2 + math.Sqrt(float64(n/2)))
	ip := make([]int, ipLen)
	wLen := n / 2
	w := make([]float64, wLen)
	fftsg.Rdft(n, 1, a, ip, w)
	fmt.Println(a)
}
```
