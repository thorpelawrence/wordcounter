package main

import "testing"

func BenchmarkNormaliseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormaliseString(`"Hello123"`)
	}
}
