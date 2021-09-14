package tokyo_go_sdk

import (
	"math"
	"testing"
)

func TestDistanceBetweenTwoPoints(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test",
			args: args{
				-7.0, 11.0, 5.0, 6.0,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistanceBetweenTwoPoints(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("DistanceBetweenTwoPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadianBetweenPoint1ToPoint2(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 0",
			args: args{
				100, 100, 200, 100,
			},
			want: 0,
		},
		{
			name: "test 45*",
			args: args{
				100, 100, 200, 200,
			},
			want: math.Pi / 4,
		},
		{
			name: "test 90*",
			args: args{
				100, 100, 100, 200,
			},
			want: math.Pi / 2,
		},
		{
			name: "test 135*",
			args: args{
				100, 100, 0, 200,
			},
			want: math.Pi * 3 / 4,
		},
		{
			name: "test 180*",
			args: args{
				100, 100, 0, 100,
			},
			want: math.Pi,
		},
		{
			name: "test 225*",
			args: args{
				100, 100, 0, 0,
			},
			want: math.Pi * 5 / 4,
		},
		{
			name: "test 270*",
			args: args{
				100, 100, 100, 0,
			},
			want: math.Pi * 3 / 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const float64EqualityThreshold = 1e-9
			if got := RadianBetweenPoint1ToPoint2(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); math.Abs(got-tt.want) > float64EqualityThreshold {
				t.Errorf("RadianBetweenPoint1ToPoint2() = %v, want %v", got, tt.want)
			}
		})
	}
}
