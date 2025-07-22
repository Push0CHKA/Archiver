package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodeTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "base tree test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &DecodingTree{
					One: &DecodingTree{
						Value: "a",
					},
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "b",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodingTable.DecodeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
