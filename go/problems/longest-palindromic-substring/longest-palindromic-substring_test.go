package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{""}, ""},
		{"1", args{"b"}, "b"},
		{"2", args{"ba"}, "b"},
		{"2e", args{"bb"}, "bb"},
		{"3", args{"bbb"}, "bbb"},
		{"simple", args{"babad"}, "bab"},
		{"simple2", args{"cbbd"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
