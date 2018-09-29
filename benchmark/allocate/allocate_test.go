package allocate

import "testing"

var loop = 1000000

func BenchmarkAllocateSliceSize1(b *testing.B) {

	for i := 1; i < b.N; i++ {
		AllocateSlice(1, loop)
	}
}
func BenchmarkAllocateSliceSize10(b *testing.B) {

	for i := 1; i < b.N; i++ {
		AllocateSlice(4000000, loop)
	}
}
