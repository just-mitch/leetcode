package linkedlist

import (
	"fmt"
	"testing"
)

func TestListToString(t *testing.T) {

	type args struct {
		l *LinkedList[int]
	}
	tests := []struct {
		args args
		want string
	}{
		{args{
			&LinkedList[int]{},
		}, ""},
		{args{
			&LinkedList[int]{
				&Node[int]{1, &Node[int]{2, &Node[int]{3, nil}}},
			},
		}, "1 2 3 "},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("to-string-%v", i), func(t *testing.T) {
			if got := fmt.Sprint(tt.args.l.ToString()); got != tt.want {
				t.Errorf("%T, toString() = %v, want %v", tt.args.l, got, tt.want)
			}
		})
	}
}
