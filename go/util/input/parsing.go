package input

import (
	"aoc-2021/util/math/vector"
	"strconv"
	"strings"
)

func BlankLineSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a double newline (ie blank line)
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}

func ParseInts(parts []string) []int {
	numbers := make([]int, len(parts))
	for i, n := range parts {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}

func SplitInts(in string, sep string) []int {
	return ParseInts(strings.Split(in, sep))
}

func ParseVec3(in string) vector.Vec3 {
	nums := SplitInts(in, ",")
	return vector.Vec3{nums[0], nums[1], nums[2]}
}
