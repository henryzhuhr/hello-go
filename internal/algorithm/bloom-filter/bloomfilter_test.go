package bloomfilter

import (
	"encoding/binary"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	tests := []struct {
		name   string
		n      uint64
		p      float64
		keyMap func(string) []byte
		want   bool // true if not nil
	}{
		{
			name: "valid",
			n:    1000,
			p:    0.01,
			keyMap: func(s string) []byte {
				return []byte(s)
			},
			want: true,
		},
		{
			name: "invalid n",
			n:    0,
			p:    0.01,
			keyMap: func(s string) []byte {
				return []byte(s)
			},
			want: false,
		},
		{
			name: "invalid p <= 0",
			n:    1000,
			p:    0,
			keyMap: func(s string) []byte {
				return []byte(s)
			},
			want: false,
		},
		{
			name: "invalid p >= 1",
			n:    1000,
			p:    1,
			keyMap: func(s string) []byte {
				return []byte(s)
			},
			want: false,
		},
		{
			name:   "nil keyMap",
			n:      1000,
			p:      0.01,
			keyMap: nil,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBloomFilter(tt.n, tt.p, tt.keyMap)
			if (got != nil) != tt.want {
				t.Errorf("NewBloomFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomFilter_String(t *testing.T) {
	bf := NewBloomFilter(100, 0.01, func(s string) []byte {
		return []byte(s)
	})

	if bf == nil {
		t.Fatal("NewBloomFilter returned nil")
	}

	// Test Add and Contains
	items := []string{"apple", "banana", "cherry"}
	for _, item := range items {
		bf.Add(item)
	}

	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("Contains(%q) = false, want true", item)
		}
	}

	// Test non-existing items
	nonExisting := []string{"dog", "cat", "elephant"}
	for _, item := range nonExisting {
		if bf.Contains(item) {
			// Note: This *could* be a false positive, but with n=100, p=0.01 and only 3 items added,
			// the probability is extremely low.
			t.Logf("Contains(%q) = true (possible false positive)", item)
		}
	}
}

func TestBloomFilter_Int(t *testing.T) {
	bf := NewBloomFilter(100, 0.01, func(i int) []byte {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(i))
		return b
	})

	if bf == nil {
		t.Fatal("NewBloomFilter returned nil")
	}

	bf.Add(42)
	bf.Add(100)

	if !bf.Contains(42) {
		t.Errorf("Contains(42) = false, want true")
	}
	if !bf.Contains(100) {
		t.Errorf("Contains(100) = false, want true")
	}
	if bf.Contains(999) {
		t.Logf("Contains(999) = true (possible false positive)")
	}
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	n := uint64(10000)
	p := 0.01
	bf := NewBloomFilter(n, p, func(i int) []byte {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(i))
		return b
	})

	// Add n items
	for i := 0; i < int(n); i++ {
		bf.Add(i)
	}

	// Check n items (should all be present)
	for i := 0; i < int(n); i++ {
		if !bf.Contains(i) {
			t.Fatalf("Contains(%d) = false, want true", i)
		}
	}

	// Check 10000 non-existing items
	falsePositives := 0
	testCount := 10000
	for i := int(n); i < int(n)+testCount; i++ {
		if bf.Contains(i) {
			falsePositives++
		}
	}

	rate := float64(falsePositives) / float64(testCount)
	t.Logf("False positive rate: %f (expected ~%f)", rate, p)

	// Allow some margin of error, e.g., 2x expected rate
	// Note: Probabilistic structures can sometimes exceed expectations slightly
	if rate > p*2 {
		t.Errorf("False positive rate %f is too high (expected %f)", rate, p)
	}
}
