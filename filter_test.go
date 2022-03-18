package tower

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		fn func(int) bool
		a  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				fn: func(a int) bool {
					return a%2 == 0
				},
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.fn, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
