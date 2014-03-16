package asyncqueue

import (
	"testing"
)

func TestSimple(t *testing.T) {
	const N = 100

	q := NewQueue()

	for i := 0; i < N; i++ {
		q.Push(i)
	}

	for want := 0; want < N; want++ {
		got := q.Pop().(int)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestConc(t *testing.T) {
	const (
		N        = 10000
		NTHREADS = 50
	)

	q := NewQueue()

	pushElems := func(id int) {
		n := N / NTHREADS
		start := id * n
		for i := start; i < start+n; i++ {
			q.Push(i)
		}
	}

	for i := 0; i < NTHREADS; i++ {
		go pushElems(i)
	}

	results := make([]int, N)

	for i := 0; i < N; i++ {
		got := q.Pop().(int)
		results[got] = got
	}

	for want := 0; want < N; want++ {
		got := results[want]
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func BenchmarkChan(b *testing.B) {
	q := NewQueue()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}
