package dtgh

import "testing"

func TestKnightProbability(t *testing.T)  {
	//t.Log(knightProbability(3, 2, 0, 0))
	t.Log(knightProbability(20, 4, 10, 10))
}

func TestKnightProbabilityIteration(t *testing.T)  {
	t.Log(knightProbabilityIteration(20, 10, 10, 10))
}

func BenchmarkKnightProbabilityIteration(b *testing.B)  {
	for i:=0; i < b.N;i++{
		knightProbabilityIteration(20, i, 10, 10)
	}
}

func BenchmarkKnightProbability(b *testing.B)  {
	for i:=0; i < b.N;i++{
		knightProbability(20, i, 10, 10)
	}
}