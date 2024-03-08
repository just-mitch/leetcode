package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]string{"fooo", "for"}}, "fo"},
		{"2", args{[]string{"fooo", "fo"}}, "fo"},
		{"no match", args{[]string{"ooo", "fo"}}, ""},
		{"empty string", args{[]string{"ooo", ""}}, ""},
		{"empty string first", args{[]string{"", "aasd"}}, ""},
		{"asdf", args{[]string{"asdf", "asdf", "asdf"}}, "asdf"},
		{"asdfasdf", args{[]string{"asdfasdf", "asdf", "asdf"}}, "asdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
