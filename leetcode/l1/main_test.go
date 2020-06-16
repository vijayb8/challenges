package l1

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test example",
			args: args{
				nums: []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("main() = %v, want %v", got, tt.want)
			}
		})
	}
}