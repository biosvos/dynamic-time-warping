package dynamic_time_warping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicTimeWarping(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			args: args{
				a: []int{1, 2, 3},
				b: []int{2, 2, 2, 3, 4},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DynamicTimeWarping(tt.args.a, tt.args.b), "DynamicTimeWarping(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
