package mathhelper

import "math"

// Does not work for MinInt** (because MinInt** * -1 = MinInt**)
func AbsInt64(i int64) int64 {
	return (-2*NegativeInt64(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt64(i int64) int64 {
	return math.MaxInt64*EqualInt64(i, math.MinInt64) + AbsInt64(i)*NotEqualInt64(i, math.MinInt64)

	//if i == math.MinInt64 {
	//	return math.MaxInt64
	//}
	//return AbsInt64(i)
}

func AntiAbsInt64(i int64) int64 {
	//return (-2*PositiveInt64(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

func DivideInt64(a, b int64) (c int64) {
	c = a / b
	delta := AntiAbsInt64(a % b)
	if b < 0 && delta < (b+1)/2 {
		if a > 0 {
			c--
			return
		}
		c++
		return
	}
	if b > 0 && delta < (-b+1)/2 {
		if a < 0 {
			c--
			return
		}
		c++
		return
	}
	return

	//return a/b + LessInt64(AntiAbsInt64(a%b), (AntiAbsInt64(b)+1)/2)*(SameSignInt64(a, b)*2-1)

	//c = a / b
	//delta := a % b
	//if delta > 0 { // To negative because |MinInt|>MaxInt
	//	delta = -delta
	//}
	//if b < 0 && delta < (b+1)/2 {
	//	if a > 0 {
	//		c--
	//		return
	//	}
	//	c++
	//	return
	//}
	//if b > 0 && delta < (-b+1)/2 {
	//	if a < 0 {
	//		c--
	//		return
	//	}
	//	c++
	//	return
	//}
	//return
}

func DivideFixInt64(a, b int64) int64 {
	if a == math.MinInt64 && b == -1 {
		return math.MaxInt64
	}
	return DivideInt64(a, b)
}
