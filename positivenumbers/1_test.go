package positivenumbers

import "testing"

func TestPositivenum(t *testing.T) {
	ints := [...]int{-1, 2, 1, 0, 3, 77, 100, 0, 0, -100, -2, 4, 1, 2, 100}
	t.Logf("ints %v", Positivenum(ints[:]))

	ints2 := [...]int{0, -1}
	t.Logf("ints %v", Positivenum(ints2[:]))
}
