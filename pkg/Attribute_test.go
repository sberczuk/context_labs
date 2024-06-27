package pkg

import (
	"cmp"
	"testing"
)

func TestBaseAttribute_LessThan(t *testing.T) {
	type args struct {
		v any
	}
	type testCase[T cmp.Ordered] struct {
		name string
		a    BaseAttribute[T]
		args args
		want bool
	}
	tests := []testCase[float64]{
		{
			name: "Pass",
			a: BaseAttribute[float64]{
				Name:  Price,
				Value: float64(200),
			},
			args: args{
				v: float64(300),
			},
			want: true,
		},
		{
			name: "Fail",
			a: BaseAttribute[float64]{
				Name:  Price,
				Value: 200,
			},
			args: args{
				v: float64(100),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.LessThan(tt.args.v); got != tt.want {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
