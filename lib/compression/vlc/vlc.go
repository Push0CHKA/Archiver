package vlc

import (
	"archiver/lib/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
	"unicode"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

// encode string to hex chunk of strings
func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)

	encoded := encodeBin(str, tbl)

	return buildEncodedFile(tbl, encoded)
}

// decode encoded string
func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)

	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBinaryCount = 4
		dataSizeBinaryCount  = 4
	)
	tableSizeBinary, data := data[:tableSizeBinaryCount], data[tableSizeBinaryCount:]
	dataSizeBinary, data := data[:dataSizeBinaryCount], data[dataSizeBinaryCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)

	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize]
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer

	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)

	buf.Write([]byte(splitByChunks(data, chunkSize).Bytes()))

	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can`t serialize table: ", err)
	}

	return tableBuf.Bytes()
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can`t decode table: ", err)
	}

	return tbl
}

// encodeBin encodes str into binary codes without spaces
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}

// bin return character code from vlc
func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

// exportText is opposite to prepareText, it prepares decoded text to export:
// it changes: ! + <lower case letter> -> to upper case letter.
//
//	i.g.: !my name is !nikita -> My name is Nikita.
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool
	for _, ch := range str {
		if ch == '!' {
			isCapital = true
			continue
		}

		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}
