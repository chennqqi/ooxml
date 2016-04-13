/*
the file is copy from https://github.com/seb-nyberg/go-docxtotex
thanks seb-nyberg

*/

package ooxml

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

type WordDocument struct {
	XMLName xml.Name `xml:" document"`
	Body    WordBody `xml:"body"`
	W       string   `xml:"w,attr"`
	Wpc     string   `xml:"wpc,attr"`
	Mc      string   `xml:"mc,attr"`
	O       string   `xml:"o,attr"`
	R       string   `xml:"r,attr"`
	M       string   `xml:"m,attr"`
	V       string   `xml:"v,attr"`
	Wp14    string   `xml:"wp14,attr"`
	Wp      string   `xml:"wp,attr"`
	W14     string   `xml:"w14,attr"`
	W10     string   `xml:"w10,attr"`
	W15     string   `xml:"w15,attr"`
	Wpg     string   `xml:"wpg,attr"`
	Wpi     string   `xml:"wpi,attr"`
	Wne     string   `xml:"wne,attr"`
	Wps     string   `xml:"wps,attr"`
}

type WordBody struct {
	XMLName    xml.Name        `xml:"body"`
	Paragraphs []WordParagraph `xml:"p"`
	WordSect   WordSectPr      `xml:"sectPr"`
}

type WordParagraph struct {
	XMLName       xml.Name          `xml:"p"`
	PPR           WordPPR           `xml:"pPr"`
	R             []WordR           `xml:"r"`
	BookmarkStart WordBookmarkStart `xml:"bookmarkStart"`
	BookmarkEnd   WordBookmarkEnd   `xml:"bookmarkEnd"`
}

//`w:color` element, specifying the color of a font and perhaps other
//  objects.
type WordColor struct {
	Val        string `xml:"val,attr"`
	ThemeColor string `xml:"val,themeColor"`
}

//``<w:rFonts>`` element, specifying typeface name for the various language
//types.

type WordFonts struct {
	Ascii string `xml:"val,ascii"`
	HAnsi string `xml:"val,hAnsi"`
}

type WordPPR struct {
	XMLName xml.Name `xml:"pPr"`
	//PStyle  WordPStyle `xml:"pStyle"`
	//NumPR   WordNumPR  `xml:"numPr"`
	RPr WordRPR `xml:"rPr"`
}
type WordPStyle struct {
	XMLName xml.Name `xml:"pStyle"`
	Value   string   `xml:"val,attr"`
}
type WordNumPR struct {
	XMLName xml.Name  `xml:"numPr"`
	NumID   WordNumID `xml:"numId"`
}
type WordNumID struct {
	XMLName xml.Name `xml:"numId"`
	Value   string   `xml:"val,attr"`
}

type WordR struct {
	XMLName xml.Name `xml:"r"`
	RPR     WordRPR  `xml:rPr"`
	T       []string `xml:"t"`
	BR      []WordBR `xml:"br"`
	CR      []WordCR `xml:"cr"`
	Tab     []WordCR `xml:"tab"`
	Drawing []WordCR `xml:"drawing"`
}
type WordBR struct {
	XMLName xml.Name `xml:"br"`
	Type    string   `xml:"type,attr"`
	Clear   string   `xml:"clear,attr"`
}

type WordDrawing struct {
	XMLName xml.Name `xml:"Drawing"`
}

type WordCR struct {
	XMLName xml.Name `xml:"cr"`
}
type WordRPR struct {
	XMLName xml.Name   `xml:"rPr"`
	I       WordI      `xml:"i"`
	U       WordU      `xml:"u"`
	B       WordB      `xml:"b"`
	Lang    WordLang   `xml:"lang"`
	RFonts  WordRFonts `xml:"rFonts"`
	Sz      WordSz     `xml:"sz"`
	SzCs    WordSzCs   `xml:"szCs"`
}

type WordSz struct {
	XMLName xml.Name `xml:"sz"`
	Val     string   `xml:"val,attr"`
}

type WordSzCs struct {
	XMLName xml.Name `xml:"szCs"`
	Val     string   `xml:"val,attr"`
}

type WordRFonts struct {
	XMLName xml.Name `xml:"rFonts"`
	Hint    string   `xml:"hint,attr"`
}

type WordLang struct {
	XMLName xml.Name `xml:"lang"`
	Val     string   `xml:"val,attr"`
}

type WordI struct {
	XMLName xml.Name `xml:"i"`
	Value   string   `xml:"val,attr"`
}
type WordB struct {
	XMLName xml.Name `xml:"b"`
	Value   string   `xml:"val,attr"`
}
type WordU struct {
	XMLName xml.Name `xml:"u"`
	Value   string   `xml:"val,attr"`
}

type WordSectPr struct {
	XMLName xml.Name    `xml:"sectPr"`
	PgSz    WordPgSz    `xml:"pgSz"`
	PgMar   WordPgMar   `xml:"pgMar"`
	Cols    WordCols    `xml:"cols"`
	DocGrid WordDocGrid `xml:"docGrid"`
}

type WordPgSz struct {
	XMLName xml.Name `xml:"pgSz"`
	W       string   `xml:"w,attr"`
	H       string   `xml:"h,attr"`
}

/*
   ``<w:pgMar>`` element, defining page margins.
*/
type WordPgMar struct {
	XMLName xml.Name `xml:"pgMar"`
	Top     string   `xml:"top,attr"`
	Right   string   `xml:"right,attr"`
	Bottom  string   `xml:"bottom,attr"`
	Left    string   `xml:"left,attr"`
	Header  string   `xml:"header,attr"`
	Footer  string   `xml:"footer,attr"`
	Gutter  string   `xml:"gutter,attr"`
}

type WordCols struct {
	XMLName xml.Name `xml:"cols"`
	Space   string   `xml:"space,attr"`
	num     string   `xml:"num,attr"`
}

type WordDocGrid struct {
	XMLName   xml.Name `xml:"docGrid"`
	Type      string   `xml:"type,attr"`
	LinePitch string   `xml:"linePitch,attr"`
	CharSpace string   `xml:"charSpace,attr"`
}

type WordBookmarkStart struct {
	XMLName xml.Name `xml:"bookmarkStart"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type WordBookmarkEnd struct {
	XMLName xml.Name `xml:"bookmarkEnd"`
	Id      string   `xml:"id,attr"`
}

func DecodeWordXML(xmlDoc []byte) (WordDocument, error) {
	var d WordDocument

	err := xml.Unmarshal(xmlDoc, &d)
	if err != nil {
		log.Fatal(err)
	}

	return d, nil
}

//add w: to every line
func EncodeWord(v interface{}, prefix string) string {

	var h xml.Encoder
	e := &xml.Encoder{printer{Writer: bufio.NewWriter(os.Stdout), prefix: "hah"}}

	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	b, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	var buffer bytes.Buffer

	for _, ln := range lines {
		//buffer.WriteString(strings.Replace(ln,"<","<",1))
		//str1:=strings.Replace(ln,"</","</w:",-1)
		//str2:=strings.Replace(str1," <"," <w:",-1)

		buffer.WriteString(ln)

		buffer.WriteString("\n")

	}

	return buffer.String()
}
