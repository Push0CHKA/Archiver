package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base name",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		hchs BinaryChunks
		want string
	}{
		{
			name: "base test",
			hchs: BinaryChunks{"01001111", "10000000"},
			want: "0100111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hchs.Join(); got != tt.want {
				t.Errorf("HexChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "base test",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
