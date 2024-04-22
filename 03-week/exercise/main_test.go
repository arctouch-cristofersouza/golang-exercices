package main

import "testing"

func Test_calculateVolume(t *testing.T) {
	type args struct {
		shapes shapes
	}
	tests := []struct {
		name string
		args args
		want float64
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateVolume(tt.args.shapes); got != tt.want {
				t.Errorf("calculateVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
