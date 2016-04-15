/*
the file is copy from https://github.com/seb-nyberg/go-docxtotex
thanks seb-nyberg

*/

package ooxml

import (
	"bufio"

	"fmt"
	"io/ioutil"
	"log"
	"os"

	"encoding/xml"
)

type mystring string

var _ xml.MarshalerAttr = (*mystring)(nil)

func (this *mystring) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	fmt.Printf("MarshalXMLAttr|%s:%s=>%s\n", n.Space, n.Local, string(*this))
	return xml.Attr{n, string(*this)}, nil
}

func (m *mystring) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	e.EncodeToken(start)
	e.EncodeToken(xml.CharData([]byte("haha")))
	e.EncodeToken(xml.EndElement{start.Name})
	return nil
}

//add w: to every line
func EncodeWord(v interface{}, filename string) string {
	f, err := os.Create(filename)

	w := bufio.NewWriter(f)

	if err != nil {
		log.Fatal(err)
	}

	enc := xml.NewEncoder(w)
	enc.Indent(" ", " ")
	//enc.SetNS()
	if err := enc.Encode(v); err != nil {
		log.Fatal(err)
	}
	f.Sync()
	w.Flush()

	return ""

}

func DecodeWord(filename string) *CT_Document {

	content, err := ioutil.ReadFile(filename)
	var v CT_Document
	//fmt.Printf("%s\n", string(content))

	err = xml.Unmarshal(content, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
	return &v

}
