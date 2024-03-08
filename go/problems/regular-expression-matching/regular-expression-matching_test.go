package regularexpressionmatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{"aa", "aa"}, true},
		{"1", args{"aa", "a"}, false},
		{"2", args{"aa", "a*"}, true},
		{"3", args{"baa", "a*"}, false},
		{"4", args{"ab", ".*"}, true},
		{"5", args{"ab", "ac*b"}, true},
		{"6", args{"ab", ".*c"}, false},
		{"7", args{"aaa", "a*a"}, true},
		{"8", args{"mississippi", "mis*is*ip*."}, true},
		{"9", args{"aaa", "ab*a*c*a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
