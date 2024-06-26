package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected :="aaaaaaaa"

	if repeated != expected {
	t.Errorf("expected %q but got %q", expected, repeated)
}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("b", 7)
	fmt.Println(res)
	// Output: bbbbbbb 
}