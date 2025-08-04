package shennonfano

import (
	"reflect"
	"testing"
)

func Test_bestDevidePosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one element",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two element",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "three element",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many element",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDevidePosition(tt.codes); got != tt.want {
				t.Errorf("bestDevidePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements, certain position",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
		{
			name: "three elements, uncertain position",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("assignCodes() = %v, want %v", tt.codes, tt.want)
			}
		})
	}
}
