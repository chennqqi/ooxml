package ooxml

import (
	"fmt"
	//"fmt"
	"testing"
)

func TestParserxml(t *testing.T) {

}

var (
	word *WordDocument
)

func TestDecodeWord(t *testing.T) {
	word = DecodeWord("./test/document.xml")

}

func TestEncodeWord(t *testing.T) {
	rst := EncodeWord(word, "")
	fmt.Printf("%s\n", rst)

}
