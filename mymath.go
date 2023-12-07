package mymath

import (
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
	return Float64frombits(Float64bits(x) &^ (1 << 63))
}

func Yn(n int, x float64) float64 {
	const Two302 = 1 << 302 // 2**302 0x52D0000000000000
	// special cases
	switch {
	case x < 0 || IsNaN(x):
		return NaN()
	case IsInf(x, 1):
		return 0
	}

	if n == 0 {
		return Y0(x)
	}
	if x == 0 {
		if n < 0 && n&1 == 1 {
			return Inf(1)
		}
		return Inf(-1)
	}
	sign := false
	if n < 0 {
		n = -n
		if n&1 == 1 {
			sign = true // sign true if n < 0 && |n| odd
		}
	}
	if n == 1 {
		if sign {
			return -Y1(x)
		}
		return Y1(x)
	}
	var b float64
	if x >= Two302 { // x > 2**302
		// (x >> n**2)
		//	    Jn(x) = cos(x-(2n+1)*pi/4)*sqrt(2/x*pi)
		//	    Yn(x) = sin(x-(2n+1)*pi/4)*sqrt(2/x*pi)
		//	    Let s=sin(x), c=cos(x),
		//		xn=x-(2n+1)*pi/4, sqt2 = sqrt(2),then
		//
		//		   n	sin(xn)*sqt2	cos(xn)*sqt2
		//		----------------------------------
		//		   0	 s-c		 c+s
		//		   1	-s-c 		-c+s
		//		   2	-s+c		-c-s
		//		   3	 s+c		 c-s

		var temp float64
		switch s, c := Sincos(x); n & 3 {
		case 0:
			temp = s - c
		case 1:
			temp = -s - c
		case 2:
			temp = -s + c
		case 3:
			temp = s + c
		}
		b = (1 / SqrtPi) * temp / Sqrt(x)
	} else {
		a := Y0(x)
		b = Y1(x)
		// quit if b is -inf
		for i := 1; i < n && !IsInf(b, -1); i++ {
			a, b = b, (float64(i+i)/x)*b-a
		}
	}
	if sign {
		return -b
	}
	return b
}
