package arithmetic

import (
	"fmt"
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
			fmt.Println(ret)
		} else {
			t.Errorf("in1:%v, in2:%v, expected: %v, output:%v", tt.in1, tt.in2, ret, tt.expected)
		}
	}
}