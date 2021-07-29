package numwordsconv

import (
	"testing"
)

func TestIntegerToWord(t *testing.T) {
	for _, tc := range testCases {
		got := IntegerToWords(tc.in)
		if tc.want != got {
			t.Errorf("IntegerToWord(%v) == %v, want %v", tc.in, got, tc.want)
		}
	}
}
