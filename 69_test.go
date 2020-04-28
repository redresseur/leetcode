package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	//t.Log(mySqrt(2))
	for i := 0; i <= 10; i++ {
		t.Log(i, mySqrt(i), mySqrt2(i))
	}
}

func TestSqrtFloat(t *testing.T) {
	for i := 100; i <= 200; i++ {
		a := math.Sqrt(float64(i))
		b := mySqrtFloat2(float64(i))
		if int(a*100000) != int(b*100000) {
			t.Fatal(i, a, b)
		}else {
			t.Log(i, a, b)
		}
	}

	//for i := 100000; i <= 200000; i++ {
	//	a := math.Sqrt(float64(i)/100000)
	//	b := mySqrtFloat2(float64(i)/100000)
	//	if int(a*100000) != int(b*100000) {
	//		t.Fatal(float64(i)/100000, a, b)
	//	}
	//}
	//mySqrtFloat2(10)
}

func TestSqrtFloat1(t *testing.T) {
	i := 0.01
	t.Log(i, math.Sqrt(float64(i)), mySqrtFloat2(float64(i)))
}


func BenchmarkMySqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrt(i)
	}
}

func BenchmarkMySqrt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrt2(i)
	}
}

func BenchmarkMySqrtFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrtFloat(float64(i))
	}
}

func BenchmarkMySqrtFloat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySqrtFloat2(float64(i))
	}
}
