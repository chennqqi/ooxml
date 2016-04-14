package ooxml

import (
	"log"
	"fmt"
	//"fmt"
	"testing"
	"github.com/juju/xml"
)

func TestParserxml(t *testing.T) {

}

var (
	word *WordDocument
)

func TestDecodeWord(t *testing.T) {

	word = DecodeWord("./test/document.xml")
	fmt.Printf("%s %s\n", word.XMLName.Local, word.XMLName.Space)

}

func TestEncodeWord(t *testing.T) {
	
	//rst := EncodeWord(word, "")
	//fmt.Printf("%s\n", rst)
	b, err := xml.MarshalIndent(word, " ", " ")
	if err != nil {
		
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)

}
