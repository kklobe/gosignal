# fftsgh

fftsgh provides a pure Go port of [Ooura FFT's](http://www.kurims.kyoto-u.ac.jp/~ooura/fft.html) fftsg_h_.c. The API is almost identical to the Ooura C version for ease of porting existing code that consumes it, which was the primary motivation. 

Some reasons:

- Pure Go for easy cross-compiling
- Unrestricted license
- Good performance

# Notes

I've only created tests and benchmarks for basic Real FFT usage, since that's 100% of my usage. However, I've no reason to believe the other functions won't be identical to the C behavior.

There are extensive comments in the fftsgh.go file from the original detailing usage. It has a more convenient interface than fftsg, but slightly lower performance.

## Installation and Usage

```$ go get -u github.com/kklobe/gosignal/fftsgh```

```
package main

import (
	"fmt"
	"github.com/kklobe/gosignal/fftsgh"
)

func main() {
	a := []float64{1, 0, 0, 0, 0, 0, 0, 0}
	n := len(a)
	fftsgh.Rdft(n, 1, a)
	fmt.Println(a)
}
```
