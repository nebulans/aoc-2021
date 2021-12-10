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
			l.Process()
			if l.corrupt != tt.corrupt {
				t.Errorf("got %v, want %v", l.corrupt, tt.corrupt)
			}
			if l.corruptChar != tt.char {
				t.Errorf("got %v, want %v", l.corruptChar, tt.char)
			}
		})
	}
}
