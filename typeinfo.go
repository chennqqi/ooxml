package ooxml

import "github.com/wxd237/ooxml/xml"

//cur hand rFonts
type WordRFonts struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rFonts"`
	Hint    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hint,attr"`
}

//cp:coreProperties
type CT_CoreProperties struct {
}

//w:document
type CT_Document struct {
	Body CT_Body `xml:"body"`
}

//w:body
type CT_Body struct {
	P      CT_P      `xml:"p"`
	Tbl    CT_Tbl    `xml:"tbl"`
	SectPr CT_SectPr `xml:"sectpr"`
}

//w:p
type CT_P struct {
	PPR CT_PPr `xml:"pPr"`
	R   CT_R   `xml:"r"`
}

//w:lvlOverride
type CT_NumLvl struct {
}

//w:num
type CT_Num struct {
}

//w:numPr
type CT_NumPr struct {
}

//w:numbering
type CT_Numbering struct {
}

//w:pgMar
type CT_PageMar struct {
}

//w:pgSz
type CT_PageSz struct {
}

//w:sectPr
type CT_SectPr struct {
}

//w:type
type CT_SectType struct {
}

//a:blip
type CT_Blip struct {
}

//a:graphic
type CT_GraphicalObject struct {
}

//a:graphicData
type CT_GraphicalObjectData struct {
}

//a:off
type CT_Point2D struct {
}

//a:xfrm
type CT_Transform2D struct {
}

//pic:blipFill
type CT_BlipFillProperties struct {
}

//pic:cNvPr
type CT_NonVisualDrawingProps struct {
}

//pic:nvPicPr
type CT_PictureNonVisual struct {
}

//pic:pic
type CT_Picture struct {
}

//pic:spPr
type CT_ShapeProperties struct {
}

//wp:inline
type CT_Inline struct {
}

//w:latentStyles
type CT_LatentStyles struct {
}

//w:lsdException
type CT_LsdException struct {
}

//w:style
type CT_Style struct {
}

//w:styles
type CT_Styles struct {
}

//w:gridCol
type CT_TblGridCol struct {
}

//w:tbl
type CT_Tbl struct {
	CT_TblPr
	CT_TblGrid
}

//w:tblGrid
type CT_TblGrid struct {
}

//w:tblLayout
type CT_TblLayoutType struct {
}

//w:tblPr
type CT_TblPr struct {
}

//w:tc
type CT_Tc struct {
}

//w:tcPr
type CT_TcPr struct {
}

//w:tcW
type CT_TblWidth struct {
}

//w:tr
type CT_Row struct {
}

//w:vMerge
type CT_VMerge struct {
}

//w:color
type CT_Color struct {
}

//w:rFonts
type CT_Fonts struct {
}

//w:rPr
type CT_RPr struct {
}

//w:sz
type CT_HpsMeasure struct {
}

//w:u
type CT_Underline struct {
}

//w:vertAlign
type CT_VerticalAlignRun struct {
}

//w:ind
type CT_Ind struct {
}

//w:jc
type CT_Jc struct {
}

//w:pPr
type CT_PPr struct {
}

//w:spacing
type CT_Spacing struct {
}

//w:br
type CT_Br struct {
}

//w:r
type CT_R struct {
}

//w:t
type CT_Text struct {
}
