package botgo

import (
	"testing"
	"time"
)

func Test_calcInterval(t *testing.T) {
	type args struct {
		maxConcurrency uint32
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"c1", args{maxConcurrency: 1}, 5 * time.Second},
		{"c3", args{maxConcurrency: 3}, 2 * time.Second},
		{"c5", args{maxConcurrency: 5}, 1 * time.Second},
		{"c10", args{maxConcurrency: 10}, 1 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcInterval(tt.args.maxConcurrency); got != tt.want {
				t.Errorf("CalcInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
