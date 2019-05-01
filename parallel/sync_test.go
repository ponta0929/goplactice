package parallel_test

import (
	"testing"

	"github.com/ponta0929/goplactice/parallel"
)

func TestWaitGroupSample(t *testing.T) {
	parallel.WaitGroupSample()
}

func TestWaitGroupRoop(t *testing.T) {
	parallel.WaitGroupRoop()
}

func TestMutexSample(t *testing.T) {
	parallel.MutexSample()
}

func TestRWMutexSample(t *testing.T) {
	parallel.RWMutexSample()
}

func TestCondSample(t *testing.T) {
	parallel.CondSample()
}

func TestCondBroadcastSample(t *testing.T) {
	parallel.CondBroadcastSample()
}

func TestOnceSample(t *testing.T) {
	parallel.OnceSample()
}

func TestOnceSample2(t *testing.T) {
	parallel.OnceSample2()
}

func TestOnceSampleError(t *testing.T) {
	parallel.OnceSampleError()
}

func TestPoolSample(t *testing.T) {
	parallel.PoolSample()
}

func TestPoolSample2(t *testing.T) {
	parallel.PoolSample2()
}

//ベンチマークが上手く動かないので封印
/*func TestPoolBenchNotUsed(t *testing.T) {
	parallel.PoolBenchNotUsed()
}

func TestPoolBenchUsed(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		parallel.PoolBenchUsed()
	}()
	wg.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				b.Fatalf("cannot dial host: %v", err)
			}
			if _, err := ioutil.ReadAll(conn); err != nil {
				b.Fatalf("cannot read: %v", err)
			}
			conn.Close()
		}()
	}
}*/
