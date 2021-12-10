package day10

import (
	"fmt"
	"testing"
)

func TestSyntaxLine_Corrupt(t *testing.T) {
	var tests = []struct {
		line    string
		corrupt bool
		char    rune
	}{
		{"{}", false, ' '},
		{"{()}", false, ' '},
		{"{()]", true, ']'},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Corrupt(%s)==(%v, %v)", tt.line, tt.corrupt, tt.char)
		t.Run(testName, func(t *testing.T) {
			l := MakeLine(tt.line)
			corrupt, char := l.Corrupt()
			if corrupt != tt.corrupt {
				t.Errorf("got %v, want %v", corrupt, tt.corrupt)
			}
			if char != tt.char {
				t.Errorf("got %v, want %v", char, tt.char)
			}
		})
	}
}
