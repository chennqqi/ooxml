package ooxml

import "encoding/xml"

type CT_DecimalNumber CT_String

type CT_OnOff CT_String

type CT_String struct {
	XmlName xml.Name
	Val     string `xml:"val,attr"`
}
