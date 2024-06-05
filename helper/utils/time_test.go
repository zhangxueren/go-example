package utils

import "testing"

func Test_t_Format(t1 *testing.T) {
	type args struct {
		timestamp int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试1", args{timestamp: 1621768906}, "2021-05-23 19:21:46"},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := t{}
			if got := t.Format(tt.args.timestamp); got != tt.want {
				t1.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
