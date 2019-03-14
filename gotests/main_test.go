package main

import (
	"testing"
	"time"
)

func TestSampale(t *testing.T) {
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
			if got := Sampale(tt.args.str, tt.args.num, tt.args.t); got != tt.want {
				t.Errorf("Sampale(%v, %v, %v) = %v, want %v", tt.args.str, tt.args.num, tt.args.t, got, tt.want)
			}
		})
	}
}
