package sample

import (
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	type args struct {
		str string
		num int
		t   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sample(tt.args.str, tt.args.num, tt.args.t); got != tt.want {
				t.Errorf("Sample(%v, %v, %v) = %v, want %v", tt.args.str, tt.args.num, tt.args.t, got, tt.want)
			}
		})
	}
}
