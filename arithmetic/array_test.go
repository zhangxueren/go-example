package arithmetic

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		in1       []int
		in2      int
		expected []int
	}{
		{[]int{1,2,3,4,5,6,7,8}, 9, []int{1,8}},
		{[]int{2,3,4,5,6,7,8}, 10, []int{1,7}},
	}

	for _, tt := range tests {
		ret := twoSum(tt.in1, tt.in2)
		if ret[0] == tt.expected[0] && ret[1] == tt.expected[1] {
			t.Logf("in1:%v, in2:%v, expected: %v, output:%v", tt.in1, tt.in2, ret, tt.expected)
		} else {
			t.Errorf("in1:%v, in2:%v, expected: %v, output:%v", tt.in1, tt.in2, ret, tt.expected)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct{
		in []byte
		expected []byte
	}{
		{[]byte{'h', 'e','l', 'l', 'o'}, []byte{'o', 'l','l', 'e', 'h'}},
	}

	for _, tt := range tests {
		in := make([]byte, len(tt.in))
		_ = copy(in, tt.in)
		ReverseString(tt.in)
		if tt.in[0] != tt.expected[0] || tt.in[len(tt.in)-1] != tt.expected[len(tt.in)-1] {
			t.Errorf("in:%v, expected: %v, output:%v", in, tt.expected, tt.in)
		} else {
			t.Logf("in:%s, expected: %s, output:%s", in, tt.expected, tt.in)
		}
	}
}