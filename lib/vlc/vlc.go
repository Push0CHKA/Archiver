package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type HexChunks []HexChunk

type HexChunk string

type BinaryChunks []BinaryChunk

type BinaryChunk string

type encodingTable map[rune]string

const chunkSize = 8

// encode string to hex chunk of strings
func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.ToHex().ToString()
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

func (hchs HexChunks) ToString() string {
	const sep = " "

	switch len(hchs) {
	case 0:
		return ""
	case 1:
		return string(hchs[0])
	}

	var buf strings.Builder

	for i, hc := range hchs {
		buf.WriteString(string(hc))

		if i < len(hchs)-1 {
			buf.WriteString(sep)
		}
	}

	return buf.String()
}

func (bchs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bchs))

	for _, chunk := range bchs {
		res = append(res, chunk.ToHex())
	}

	return res
}

func (bch BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bch), 2, chunkSize)
	if err != nil {
		panic("can`t parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

// encodeBin encodes str into binary codes without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}

}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Nikita -> !my name is !nikita
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
