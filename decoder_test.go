package ooxml

import (
	"fmt"
	//"log"
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
	fmt.Printf("%+v\n", word)
	//fmt.Printf("%s %s\n", word.XMLName.Local, word.XMLName.Space)

}

func TestEncodeWord(t *testing.T) {
	rst := EncodeWord(word, "")
	fmt.Printf("%s\n", rst)

}
