package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMarshalUnmarshal(b *testing.B) {
	value := &Device{
		Type: 10,
		Vendor: &KeyValue{
			ID:   100,
			Name: "Vendor",
		},
		OS: &KeyValueVersion{
			ID:   10,
			Name: "OS",
			Version: &Version{
				Major: "OS Major",
				Minor: "OS Minor",
			},
		},
	}
	data := value.Marshal()
	if !assert.Equal(b, 60, len(data)) {
		return
	}
	value2 := &Device{}
	err := value2.Unmarshal(data)
	if err != nil {
		b.Fatal(err)
	}
	if !assert.Equal(b, value, value2) {
		return
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value.Marshal()
		value2.Unmarshal(data)
	}
}

// Size 60b
// BenchmarkMarshalUnmarshal-8   	 5000000	       267 ns/op	      88 B/op	       5 allocs/op
// BenchmarkMarshalUnmarshal-8   	 5000000	       259 ns/op	      88 B/op	       5 allocs/op
// BenchmarkMarshalUnmarshal-8   	 5000000	       269 ns/op	      88 B/op	       5 allocs/op
