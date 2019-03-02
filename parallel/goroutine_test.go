package parallel_test

import (
	"sync"
	"testing"

	"github.com/ponta0929/goplactice/parallel"
)

func TestCallGoroutine(t *testing.T) {
	parallel.CallGoroutine()
}

func TestGoClosure(t *testing.T) {
	parallel.GoClosure()
}

func TestGoClosureRangeBad(t *testing.T) {
	parallel.GoClosureRangeBad()
}

func TestGoClosureRangeGood(t *testing.T) {
	parallel.GoClosureRangeGood()
}

func TestGoroutineMemory(t *testing.T) {
	parallel.GoroutineMemory()
}

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
