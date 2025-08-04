package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

const chunkSize = 8

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, part := range data {
		res = append(res, NewBinChunk(part))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bchs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bchs))

	for _, bch := range bchs {
		res = append(res, bch.Byte())
	}

	return res
}

func (bch BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bch), 2, chunkSize)
	if err != nil {
		panic("can`t parse binary chunk: " + err.Error())
	}

	return byte(num)
}

// Join joins nhunks into one line and returns as string
func (bchs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bch := range bchs {
		buf.WriteString(string(bch))
	}

	return buf.String()
}

// splitByChunks split binary string by chunks with given,
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunkCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunkCount++
	}

	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}
