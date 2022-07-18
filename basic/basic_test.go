package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{8, 15, 17},
		{12, 35, 37},
		{3000000, 4000000, 5000000},
	}

	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("test error: actual=%d, expected=%d", actual, tt.c)
		}
	}
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		s string
		l int
	}{
		{"", 0},
		{"abc", 3},
		{"abcab", 3},
		{"ababcb", 3},
		{"今天今天不错啊", 5},
		{"今天今天", 2},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubstr(tt.s); actual != tt.l {
			t.Errorf("test error: str=%s, actual=%d, expected=%d", tt.s, actual, tt.l)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	tests := []struct {
		s string
		l int
	}{
		{"", 0},
		{"abc", 3},
		{"abcab", 3},
		{"ababcb", 3},
		{"abcabcabcd", 4},
		{"今天今天不错啊", 5},
		{"今天今天", 2},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if actual := lengthOfNonRepeatingSubstr(tt.s); actual != tt.l {
				b.Errorf("test error: str=%s, actual=%d, expected=%d", tt.s, actual, tt.l)
			}
		}
	}
}
