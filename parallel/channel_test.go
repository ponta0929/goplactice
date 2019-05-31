package parallel_test

import (
	"testing"

	"github.com/ponta0929/goplactice/parallel"
)

func TestSampleChannel(t *testing.T) {
	parallel.SampleChannel()
}

func TestDeclarationChannel(t *testing.T) {
	parallel.DeclarationChannel()
}

func BenchmarkDeclarationCannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallel.DeclarationChannel()
	}
}

func TestChannelCloseSample(t *testing.T) {
	parallel.ChannelCloseSample()
}

func TestSampleBafferChannel(t *testing.T) {
	parallel.SampleBafferChannel()
}
