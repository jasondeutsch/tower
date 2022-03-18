package tower

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		a  []int
		fn func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				a: []int{1, 2, 3},
				fn: func(a int) string {
					return strconv.Itoa(a)
				},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.a, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
