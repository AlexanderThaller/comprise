package comprise

import "testing"

func TestFormatNumberToID(t *testing.T) {
	tests := map[uint32]string{
		1234567890: "098-765-432-1",
		1000000000: "000-000-000-1",
		1:          "100-000-000-0",
		0:          "000-000-000-0",
		4294967295: "592-769-492-4",
	}

	for value, expected := range tests {
		got := formatNumberToID(value)

		if expected != got {
			t.Fatalf("got: %v, expected: %v", got, expected)
		}
	}
}

func BenchmarkFormatNumberToID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatNumberToID(0)
	}
}
