package split

import (
	"reflect"
	"testing"
)

// 测试

func TestSplit(t *testing.T) {
	got := Split("Wang爱Guan爱yu", "爱")
	want := []string{"Wang", "Guan", "yu"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("llllabbbb", "a")
	}
}
