package mathhelper

import "testing"

func TestRound(t *testing.T) {
	type testElement struct {
		f float64
		i int64
	}

	test := []testElement{
		// 0
		testElement{
			0.5,
			1,
		},

		// 1
		testElement{
			0,
			0,
		},

		// 2
		testElement{
			-0.5,
			-1,
		},

		// 3
		testElement{
			0.48999999,
			0,
		},

		// 4
		testElement{
			-0.4899999,
			0,
		},

		// 5
		testElement{
			0.599999,
			1,
		},

		// 6
		testElement{
			-0.599999,
			-1,
		},

		// 7
		testElement{
			99.999999999,
			100,
		},

		// 8
		testElement{
			99.48999999999999999999999,
			99,
		},

		// 9
		testElement{
			99.0000000001,
			99,
		},
	}
	for j, v := range test {
		i := Round(v.f)
		if i != v.i {
			t.Errorf("\nTestRound - %v.\nExpected int64: %v\ngot: %v", j, v.i, i)
		}
	}
}
