package mathhelper

/*
import "strings"

type Sign uint8

const (
	nan   = "NaN"
	delim = '.'
)
const (
	base     = 10000
	groupLen = 4 // Number of 10-based digits stored together, =lg(base)
)

const (
	Positive Sign = iota
	Negative      = iota
	NaN           = iota
)

type Numeric struct {
	sign   Sign
	digits []int16
	weight int16
}

func parseNatural(s string) (r []int16) {
	if len(s) == 0 {
		return nil
	}

	if firstLen := len(s) % groupLen; firstLen != 0 {
		s = strings.Repeat("0", groupLen-firstLen)
	}

	r = make([]int16, len(s)/groupLen)

	for i := range r {
		//k:=int16(1)
		//for j:=groupLen-1; j>=0; j--{
		//	r[i]=s[i*groupLen+j]*k
		//	k*=10
		//}
		r[i] = s[i*groupLen+0]*1000 +
			s[i*groupLen+1]*100 +
			s[i*groupLen+2]*10 +
			s[i*groupLen+3]*1
	}

	return r
}

func (n *Numeric) parseUnsigned(s string) bool {
	if len(s) == 0 { // Nothing to parse
		return false
	}

	delimPos := -1
	for i := 0; i < len(s); i++ {
		if s[i] == delim {
			if delimPos == -1 {
				delimPos = i
			} else { // Send delimiter found - error
				return nil, false
			}
		} else if s[i] < '0' || s[i] > '9' { // If char is not delimiter than it can be only 10-base digit
			return nil, false
		}
	}

	if delimPos != -1 {
		delimPos = len(s)
	}

	n.digits = parseNatural(s[:delimPos])
}

func (n *Numeric) setString(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s == nan {
		n.sign = NaN
		n.digits = nil
		n.weight = 0

		return true
	}

	switch s[0] {
	case '-':
		n.sign = Negative
		s = s[1:]
	case '+':
		n.sign = Positive
		s = s[1:]
	default:
		n.sign = Positive
	}

	return n.parseUnsigned(s)
}

func (n *Numeric) SetString(s string) (*Numeric, bool) {
	if n.setString(s) {
		return n, true
	}
	return nil, false
}
*/