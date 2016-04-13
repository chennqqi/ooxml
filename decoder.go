/*
the file is copy from https://github.com/seb-nyberg/go-docxtotex
thanks seb-nyberg

*/

package ooxml

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/wxd237/ooxml/xml"
)

// cur hand color
type WordColor struct {
	XMLName xml.Name `xml:"color"`
	Val     string   `xml:"val,attr"`
}

// cur hand sz
type WordSz struct {
	XMLName xml.Name `xml:"sz"`
	Val     string   `xml:"val,attr"`
}

// cur hand tblW
type WordTblW struct {
	XMLName xml.Name `xml:"tblW"`
	Type    string   `xml:"type,attr"`
	W       string   `xml:"w,attr"`
}

// cur hand fldChar
type WordFldChar struct {
	XMLName     xml.Name `xml:"fldChar"`
	FldCharType string   `xml:"fldCharType,attr"`
}

// cur hand cNvSpPr
type WordCNvSpPr struct {
	XMLName xml.Name `xml:"cNvSpPr"`
}

// cur hand titlePg
type WordTitlePg struct {
	XMLName xml.Name `xml:"titlePg"`
}

// cur hand shd
type WordShd struct {
	XMLName xml.Name `xml:"shd"`
	Fill    string   `xml:"fill,attr"`
	Val     string   `xml:"val,attr"`
	Color   string   `xml:"color,attr"`
}

// cur hand bookmarkStart
type WordBookmarkStart struct {
	XMLName xml.Name `xml:"bookmarkStart"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// cur hand positionV
type WordPositionV struct {
	XMLName      xml.Name      `xml:"positionV"`
	PosOffset    WordPosOffset `xml:"posOffset"`
	RelativeFrom string        `xml:"relativeFrom,attr"`
}

// cur hand pic
type WordPic struct {
	XMLName  xml.Name     `xml:"pic"`
	SpPr     WordSpPr     `xml:"spPr"`
	NvPicPr  WordNvPicPr  `xml:"nvPicPr"`
	BlipFill WordBlipFill `xml:"blipFill"`
	Pic      string       `xml:"pic,attr"`
}

// cur hand wrap
type WordWrap struct {
	XMLName xml.Name `xml:"wrap"`
	Type    string   `xml:"type,attr"`
}

// cur hand extent
type WordExtent struct {
	XMLName xml.Name `xml:"extent"`
	Cx      string   `xml:"cx,attr"`
	Cy      string   `xml:"cy,attr"`
}

// cur hand stroke
type WordStroke struct {
	XMLName xml.Name `xml:"stroke"`
	On      string   `xml:"on,attr"`
}

// cur hand instrText
type WordInstrText struct {
	XMLName xml.Name `xml:"instrText"`
	Space   string   `xml:"space,attr"`
}

// cur hand docPr
type WordDocPr struct {
	XMLName xml.Name `xml:"docPr"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// cur hand b
type WordB struct {
	XMLName xml.Name `xml:"b"`
}

// cur hand wordWrap
type WordWordWrap struct {
	XMLName xml.Name `xml:"wordWrap"`
	Val     string   `xml:"val,attr"`
}

// cur hand insideV
type WordInsideV struct {
	XMLName xml.Name `xml:"insideV"`
	Sz      string   `xml:"sz,attr"`
	Space   string   `xml:"space,attr"`
	Val     string   `xml:"val,attr"`
	Color   string   `xml:"color,attr"`
}

// cur hand wrapPolygon
type WordWrapPolygon struct {
	XMLName xml.Name     `xml:"wrapPolygon"`
	Start   WordStart    `xml:"start"`
	LineTo  []WordLineTo `xml:"lineTo"`
}

// cur hand u
type WordU struct {
	XMLName xml.Name `xml:"u"`
	Val     string   `xml:"val,attr"`
}

// cur hand xfrm
type WordXfrm struct {
	XMLName xml.Name `xml:"xfrm"`
	Off     WordOff  `xml:"off"`
	Ext     WordExt  `xml:"ext"`
}

// cur hand miter
type WordMiter struct {
	XMLName xml.Name `xml:"miter"`
}

// cur hand anchorlock
type WordAnchorlock struct {
	XMLName xml.Name `xml:"anchorlock"`
}

// cur hand tcW
type WordTcW struct {
	XMLName xml.Name `xml:"tcW"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand fillRect
type WordFillRect struct {
	XMLName xml.Name `xml:"fillRect"`
}

// cur hand anchor
type WordAnchor struct {
	XMLName      xml.Name         `xml:"anchor"`
	PositionH    WordPositionH    `xml:"positionH"`
	EffectExtent WordEffectExtent `xml:"effectExtent"`
	WrapTight    WordWrapTight    `xml:"wrapTight"`
	DocPr        WordDocPr        `xml:"docPr"`
	Graphic      WordGraphic      `xml:"graphic"`
	//SimplePos         WordSimplePos         `xml:"simplePos"`
	PositionV         WordPositionV         `xml:"positionV"`
	Extent            WordExtent            `xml:"extent"`
	CNvGraphicFramePr WordCNvGraphicFramePr `xml:"cNvGraphicFramePr"`
	Locked            string                `xml:"locked,attr"`
	LayoutInCell      string                `xml:"layoutInCell,attr"`
	AllowOverlap      string                `xml:"allowOverlap,attr"`
	SimplePos         string                `xml:"simplePos,attr"`
	BehindDoc         string                `xml:"behindDoc,attr"`
	DistL             string                `xml:"distL,attr"`
	DistR             string                `xml:"distR,attr"`
	RelativeHeight    string                `xml:"relativeHeight,attr"`
	DistT             string                `xml:"distT,attr"`
	DistB             string                `xml:"distB,attr"`
}

// cur hand i
type WordI struct {
	XMLName xml.Name `xml:"i"`
}

// cur hand Fallback
type WordFallback struct {
	XMLName xml.Name `xml:"Fallback"`
	Pict    WordPict `xml:"pict"`
}

// cur hand cNvPicPr
type WordCNvPicPr struct {
	XMLName  xml.Name     `xml:"cNvPicPr"`
	PicLocks WordPicLocks `xml:"picLocks"`
}

// cur hand pgBorders
type WordPgBorders struct {
	XMLName xml.Name   `xml:"pgBorders"`
	Top     WordTop    `xml:"top"`
	Left    WordLeft   `xml:"left"`
	Bottom  WordBottom `xml:"bottom"`
	Right   WordRight  `xml:"right"`
}

// cur hand off
type WordOff struct {
	XMLName xml.Name `xml:"off"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

// cur hand tblCellMar
type WordTblCellMar struct {
	XMLName xml.Name   `xml:"tblCellMar"`
	Top     WordTop    `xml:"top"`
	Left    WordLeft   `xml:"left"`
	Bottom  WordBottom `xml:"bottom"`
	Right   WordRight  `xml:"right"`
}

// cur hand gridCol
type WordGridCol struct {
	XMLName xml.Name `xml:"gridCol"`
	W       string   `xml:"w,attr"`
}

// cur hand tailEnd
type WordTailEnd struct {
	XMLName xml.Name `xml:"tailEnd"`
	Type    string   `xml:"type,attr"`
	W       string   `xml:"w,attr"`
	Len     string   `xml:"len,attr"`
}

// cur hand spPr
type WordSpPr struct {
	XMLName  xml.Name     `xml:"spPr"`
	Xfrm     WordXfrm     `xml:"xfrm"`
	PrstGeom WordPrstGeom `xml:"prstGeom"`
	NoFill   WordNoFill   `xml:"noFill"`
	Ln       WordLn       `xml:"ln"`
}

// cur hand bottom
type WordBottom struct {
	XMLName xml.Name `xml:"bottom"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand pgNumType
type WordPgNumType struct {
	XMLName xml.Name `xml:"pgNumType"`
	Start   string   `xml:"start,attr"`
}

// cur hand docGrid
type WordDocGrid struct {
	XMLName   xml.Name `xml:"docGrid"`
	Type      string   `xml:"type,attr"`
	LinePitch string   `xml:"linePitch,attr"`
	CharSpace string   `xml:"charSpace,attr"`
}

// cur hand inline
type WordInline struct {
	XMLName           xml.Name              `xml:"inline"`
	DocPr             WordDocPr             `xml:"docPr"`
	CNvGraphicFramePr WordCNvGraphicFramePr `xml:"cNvGraphicFramePr"`
	Graphic           WordGraphic           `xml:"graphic"`
	Extent            WordExtent            `xml:"extent"`
	EffectExtent      WordEffectExtent      `xml:"effectExtent"`
	DistT             string                `xml:"distT,attr"`
	DistB             string                `xml:"distB,attr"`
	DistL             string                `xml:"distL,attr"`
	DistR             string                `xml:"distR,attr"`
}

// cur hand bodyPr
type WordBodyPr struct {
	XMLName xml.Name `xml:"bodyPr"`
	Wrap    string   `xml:"wrap,attr"`
	LIns    string   `xml:"lIns,attr"`
	TIns    string   `xml:"tIns,attr"`
	RIns    string   `xml:"rIns,attr"`
	BIns    string   `xml:"bIns,attr"`
	Upright string   `xml:"upright,attr"`
}

// cur hand wrapTight
type WordWrapTight struct {
	XMLName     xml.Name        `xml:"wrapTight"`
	WrapPolygon WordWrapPolygon `xml:"wrapPolygon"`
	WrapText    string          `xml:"wrapText,attr"`
}

// cur hand bCs
type WordBCs struct {
	XMLName xml.Name `xml:"bCs"`
	Val     string   `xml:"val,attr"`
}

// cur hand vertAlign
type WordVertAlign struct {
	XMLName xml.Name `xml:"vertAlign"`
	Val     string   `xml:"val,attr"`
}

// cur hand rStyle
type WordRStyle struct {
	XMLName xml.Name `xml:"rStyle"`
	Val     string   `xml:"val,attr"`
}

// cur hand lineTo
type WordLineTo struct {
	XMLName xml.Name `xml:"lineTo"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

// cur hand ln
type WordLn struct {
	XMLName xml.Name   `xml:"ln"`
	NoFill  WordNoFill `xml:"noFill"`
	Miter   WordMiter  `xml:"miter"`
	W       string     `xml:"w,attr"`
}

// cur hand object
type WordObject struct {
	XMLName   xml.Name      `xml:"object"`
	OLEObject WordOLEObject `xml:"OLEObject"`
	Shape     WordShape     `xml:"shape"`
}

// cur hand fill
type WordFill struct {
	XMLName   xml.Name `xml:"fill"`
	On        string   `xml:"on,attr"`
	Focussize string   `xml:"focussize,attr"`
}

// cur hand insideH
type WordInsideH struct {
	XMLName xml.Name `xml:"insideH"`
	Val     string   `xml:"val,attr"`
	Color   string   `xml:"color,attr"`
	Sz      string   `xml:"sz,attr"`
	Space   string   `xml:"space,attr"`
}

// cur hand wsp
type WordWsp struct {
	XMLName xml.Name    `xml:"wsp"`
	CNvSpPr WordCNvSpPr `xml:"cNvSpPr"`
	SpPr    WordSpPr    `xml:"spPr"`
	Txbx    WordTxbx    `xml:"txbx"`
	BodyPr  WordBodyPr  `xml:"bodyPr"`
}

// cur hand cNvPr
type WordCNvPr struct {
	XMLName xml.Name `xml:"cNvPr"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// cur hand left
type WordLeft struct {
	XMLName xml.Name `xml:"left"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand wrapNone
type WordWrapNone struct {
	XMLName xml.Name `xml:"wrapNone"`
}

// cur hand prstDash
type WordPrstDash struct {
	XMLName xml.Name `xml:"prstDash"`
	Val     string   `xml:"val,attr"`
}

// cur hand simplePos
type WordSimplePos struct {
	XMLName xml.Name `xml:"simplePos"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

// cur hand prstGeom
type WordPrstGeom struct {
	XMLName xml.Name  `xml:"prstGeom"`
	AvLst   WordAvLst `xml:"avLst"`
	Prst    string    `xml:"prst,attr"`
}

// cur hand numPr
type WordNumPr struct {
	XMLName xml.Name  `xml:"numPr"`
	Ilvl    WordIlvl  `xml:"ilvl"`
	NumId   WordNumId `xml:"numId"`
}

// cur hand autoSpaceDE
type WordAutoSpaceDE struct {
	XMLName xml.Name `xml:"autoSpaceDE"`
	Val     string   `xml:"val,attr"`
}

// cur hand adjustRightInd
type WordAdjustRightInd struct {
	XMLName xml.Name `xml:"adjustRightInd"`
	Val     string   `xml:"val,attr"`
}

// cur hand gd
type WordGd struct {
	XMLName xml.Name `xml:"gd"`
	Name    string   `xml:"name,attr"`
	Fmla    string   `xml:"fmla,attr"`
}

// cur hand graphicFrameLocks
type WordGraphicFrameLocks struct {
	XMLName        xml.Name `xml:"graphicFrameLocks"`
	A              string   `xml:"a,attr"`
	NoChangeAspect string   `xml:"noChangeAspect,attr"`
}

// cur hand cNvGraphicFramePr
type WordCNvGraphicFramePr struct {
	XMLName           xml.Name              `xml:"cNvGraphicFramePr"`
	GraphicFrameLocks WordGraphicFrameLocks `xml:"graphicFrameLocks"`
}

// cur hand drawing
type WordDrawing struct {
	XMLName xml.Name   `xml:"drawing"`
	Anchor  WordAnchor `xml:"anchor"`
}

// cur hand szCs
type WordSzCs struct {
	XMLName xml.Name `xml:"szCs"`
	Val     string   `xml:"val,attr"`
}

// cur hand ext
type WordExt struct {
	XMLName xml.Name `xml:"ext"`
	Cx      string   `xml:"cx,attr"`
	Cy      string   `xml:"cy,attr"`
}

// cur hand tcPr
type WordTcPr struct {
	XMLName xml.Name   `xml:"tcPr"`
	VAlign  WordVAlign `xml:"vAlign"`
	TcW     WordTcW    `xml:"tcW"`
}

// cur hand solidFill
type WordSolidFill struct {
	XMLName xml.Name    `xml:"solidFill"`
	SrgbClr WordSrgbClr `xml:"srgbClr"`
}

// cur hand Choice
type WordChoice struct {
	XMLName  xml.Name    `xml:"Choice"`
	Drawing  WordDrawing `xml:"drawing"`
	Requires string      `xml:"Requires,attr"`
}

// cur hand tbl
type WordTbl struct {
	XMLName xml.Name    `xml:"tbl"`
	Tr      []WordTr    `xml:"tr"`
	TblPr   WordTblPr   `xml:"tblPr"`
	TblGrid WordTblGrid `xml:"tblGrid"`
}

// cur hand cols
type WordCols struct {
	XMLName xml.Name `xml:"cols"`
	Space   string   `xml:"space,attr"`
	Num     string   `xml:"num,attr"`
}

// cur hand document
type WordDocument struct {
	XMLName       xml.Name `xml:"document"`
	Body          WordBody `xml:"body"`
	V             string   `xml:"v,attr"`
	Ignorable     string   `xml:"Ignorable,attr"`
	Wpc           string   `xml:"wpc,attr"`
	O             string   `xml:"o,attr"`
	W14           string   `xml:"w14,attr"`
	W10           string   `xml:"w10,attr"`
	W15           string   `xml:"w15,attr"`
	Wpg           string   `xml:"wpg,attr"`
	Wne           string   `xml:"wne,attr"`
	Wps           string   `xml:"wps,attr"`
	WpsCustomData string   `xml:"wpsCustomData,attr"`
	Mc            string   `xml:"mc,attr"`
	Wp14          string   `xml:"wp14,attr"`
	Wp            string   `xml:"wp,attr"`
	W             string   `xml:"w,attr"`
	Wpi           string   `xml:"wpi,attr"`
	R             string   `xml:"r,attr"`
	M             string   `xml:"m,attr"`
}

// cur hand pPr
type WordPPr struct {
	XMLName xml.Name `xml:"pPr"`
	Ind     WordInd  `xml:"ind"`
	RPr     WordRPr  `xml:"rPr"`
}

// cur hand effectExtent
type WordEffectExtent struct {
	XMLName xml.Name `xml:"effectExtent"`
	B       string   `xml:"b,attr"`
	L       string   `xml:"l,attr"`
	T       string   `xml:"t,attr"`
	R       string   `xml:"r,attr"`
}

// cur hand shape
type WordShape struct {
	XMLName    xml.Name       `xml:"shape"`
	Fill       WordFill       `xml:"fill"`
	Stroke     WordStroke     `xml:"stroke"`
	Imagedata  WordImagedata  `xml:"imagedata"`
	Lock       WordLock       `xml:"lock"`
	Wrap       WordWrap       `xml:"wrap"`
	Anchorlock WordAnchorlock `xml:"anchorlock"`
	Path       WordPath       `xml:"path"`
	Type       string         `xml:"type,attr"`
	Style      string         `xml:"style,attr"`
	Ole        string         `xml:"ole,attr"`
	Filled     string         `xml:"filled,attr"`
	Stroked    string         `xml:"stroked,attr"`
	Coordsize  string         `xml:"coordsize,attr"`
	Id         string         `xml:"id,attr"`
	Spt        string         `xml:"spt,attr"`
}

// cur hand tc
type WordTc struct {
	XMLName xml.Name `xml:"tc"`
	TcPr    WordTcPr `xml:"tcPr"`
	P       WordP    `xml:"p"`
}

// cur hand start
type WordStart struct {
	XMLName xml.Name `xml:"start"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

// cur hand blip
type WordBlip struct {
	XMLName xml.Name `xml:"blip"`
	Embed   string   `xml:"embed,attr"`
}

// cur hand trHeight
type WordTrHeight struct {
	XMLName xml.Name `xml:"trHeight"`
	HRule   string   `xml:"hRule,attr"`
	Val     string   `xml:"val,attr"`
}

// cur hand right
type WordRight struct {
	XMLName xml.Name `xml:"right"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand gridSpan
type WordGridSpan struct {
	XMLName xml.Name `xml:"gridSpan"`
	Val     string   `xml:"val,attr"`
}

// cur hand srgbClr
type WordSrgbClr struct {
	XMLName xml.Name `xml:"srgbClr"`
	Val     string   `xml:"val,attr"`
}

// cur hand AlternateContent
type WordAlternateContent struct {
	XMLName  xml.Name     `xml:"AlternateContent"`
	Choice   WordChoice   `xml:"Choice"`
	Fallback WordFallback `xml:"Fallback"`
}

// cur hand avLst
type WordAvLst struct {
	XMLName xml.Name `xml:"avLst"`
}

// cur hand lock
type WordLock struct {
	XMLName     xml.Name `xml:"lock"`
	Text        string   `xml:"text,attr"`
	Aspectratio string   `xml:"aspectratio,attr"`
	Ext         string   `xml:"ext,attr"`
	Grouping    string   `xml:"grouping,attr"`
	Rotation    string   `xml:"rotation,attr"`
}

// cur hand textbox
type WordTextbox struct {
	XMLName     xml.Name        `xml:"textbox"`
	TxbxContent WordTxbxContent `xml:"txbxContent"`
	Inset       string          `xml:"inset,attr"`
}

// cur hand nvPicPr
type WordNvPicPr struct {
	XMLName  xml.Name     `xml:"nvPicPr"`
	CNvPr    WordCNvPr    `xml:"cNvPr"`
	CNvPicPr WordCNvPicPr `xml:"cNvPicPr"`
}

// cur hand noFill
type WordNoFill struct {
	XMLName xml.Name `xml:"noFill"`
}

// cur hand jc
type WordJc struct {
	XMLName xml.Name `xml:"jc"`
	Val     string   `xml:"val,attr"`
}

// cur hand tblGrid
type WordTblGrid struct {
	XMLName xml.Name      `xml:"tblGrid"`
	GridCol []WordGridCol `xml:"gridCol"`
}

// cur hand bookmarkEnd
type WordBookmarkEnd struct {
	XMLName xml.Name `xml:"bookmarkEnd"`
	Id      string   `xml:"id,attr"`
}

// cur hand path
type WordPath struct {
	XMLName xml.Name `xml:"path"`
}

// cur hand footerReference
type WordFooterReference struct {
	XMLName xml.Name `xml:"footerReference"`
	Type    string   `xml:"type,attr"`
	Id      string   `xml:"id,attr"`
}

// cur hand tblBorders
type WordTblBorders struct {
	XMLName xml.Name    `xml:"tblBorders"`
	Bottom  WordBottom  `xml:"bottom"`
	Right   WordRight   `xml:"right"`
	InsideH WordInsideH `xml:"insideH"`
	InsideV WordInsideV `xml:"insideV"`
	Top     WordTop     `xml:"top"`
	Left    WordLeft    `xml:"left"`
}

// cur hand rPr
type WordRPr struct {
	XMLName xml.Name   `xml:"rPr"`
	RFonts  WordRFonts `xml:"rFonts"`
}

// cur hand positionH
type WordPositionH struct {
	XMLName      xml.Name      `xml:"positionH"`
	PosOffset    WordPosOffset `xml:"posOffset"`
	RelativeFrom string        `xml:"relativeFrom,attr"`
}

// cur hand spacing
type WordSpacing struct {
	XMLName  xml.Name `xml:"spacing"`
	Line     string   `xml:"line,attr"`
	LineRule string   `xml:"lineRule,attr"`
}

// cur hand pStyle
type WordPStyle struct {
	XMLName xml.Name `xml:"pStyle"`
	Val     string   `xml:"val,attr"`
}

// cur hand ind
type WordInd struct {
	XMLName   xml.Name `xml:"ind"`
	FirstLine string   `xml:"firstLine,attr"`
}

// cur hand pgMar
type WordPgMar struct {
	XMLName xml.Name `xml:"pgMar"`
	Header  string   `xml:"header,attr"`
	Footer  string   `xml:"footer,attr"`
	Gutter  string   `xml:"gutter,attr"`
	Top     string   `xml:"top,attr"`
	Right   string   `xml:"right,attr"`
	Bottom  string   `xml:"bottom,attr"`
	Left    string   `xml:"left,attr"`
}

// cur hand txbxContent
type WordTxbxContent struct {
	XMLName xml.Name `xml:"txbxContent"`
	P       []WordP  `xml:"p"`
}

// cur hand tabs
type WordTabs struct {
	XMLName xml.Name `xml:"tabs"`
	Tab     WordTab  `xml:"tab"`
}

// cur hand tblStyle
type WordTblStyle struct {
	XMLName xml.Name `xml:"tblStyle"`
	Val     string   `xml:"val,attr"`
}

// cur hand widowControl
type WordWidowControl struct {
	XMLName xml.Name `xml:"widowControl"`
}

// cur hand pgSz
type WordPgSz struct {
	XMLName xml.Name `xml:"pgSz"`
	W       string   `xml:"w,attr"`
	H       string   `xml:"h,attr"`
}

// cur hand vAlign
type WordVAlign struct {
	XMLName xml.Name `xml:"vAlign"`
	Val     string   `xml:"val,attr"`
}

// cur hand sectPr
type WordSectPr struct {
	XMLName         xml.Name              `xml:"sectPr"`
	FooterReference []WordFooterReference `xml:"footerReference"`
	PgSz            WordPgSz              `xml:"pgSz"`
	PgMar           WordPgMar             `xml:"pgMar"`
	PgBorders       WordPgBorders         `xml:"pgBorders"`
	PgNumType       WordPgNumType         `xml:"pgNumType"`
	Cols            WordCols              `xml:"cols"`
	TitlePg         WordTitlePg           `xml:"titlePg"`
	DocGrid         WordDocGrid           `xml:"docGrid"`
}

// cur hand body
type WordBody struct {
	XMLName     xml.Name          `xml:"body"`
	P           []WordP           `xml:"p"`
	Tbl         []WordTbl         `xml:"tbl"`
	BookmarkEnd []WordBookmarkEnd `xml:"bookmarkEnd"`
	SectPr      WordSectPr        `xml:"sectPr"`
}

// cur hand lang
type WordLang struct {
	XMLName xml.Name `xml:"lang"`
}

// cur hand r
type WordR struct {
	XMLName xml.Name    `xml:"r"`
	Drawing WordDrawing `xml:"drawing"`
	RPr     WordRPr     `xml:"rPr"`
}

// cur hand OLEObject
type WordOLEObject struct {
	XMLName     xml.Name        `xml:"OLEObject"`
	LockedField WordLockedField `xml:"LockedField"`
	ObjectID    string          `xml:"ObjectID,attr"`
	Id          string          `xml:"id,attr"`
	Type        string          `xml:"Type,attr"`
	ProgID      string          `xml:"ProgID,attr"`
	ShapeID     string          `xml:"ShapeID,attr"`
	DrawAspect  string          `xml:"DrawAspect,attr"`
}

// cur hand trPr
type WordTrPr struct {
	XMLName  xml.Name     `xml:"trPr"`
	TrHeight WordTrHeight `xml:"trHeight"`
	Jc       WordJc       `xml:"jc"`
}

// cur hand t
type WordT struct {
	XMLName xml.Name `xml:"t"`
}

// cur hand tblLayout
type WordTblLayout struct {
	XMLName xml.Name `xml:"tblLayout"`
	Type    string   `xml:"type,attr"`
}

// cur hand tblPr
type WordTblPr struct {
	XMLName    xml.Name       `xml:"tblPr"`
	TblW       WordTblW       `xml:"tblW"`
	Jc         WordJc         `xml:"jc"`
	TblInd     WordTblInd     `xml:"tblInd"`
	TblLayout  WordTblLayout  `xml:"tblLayout"`
	TblCellMar WordTblCellMar `xml:"tblCellMar"`
	TblStyle   WordTblStyle   `xml:"tblStyle"`
}

// cur hand headEnd
type WordHeadEnd struct {
	XMLName xml.Name `xml:"headEnd"`
	Type    string   `xml:"type,attr"`
	W       string   `xml:"w,attr"`
	Len     string   `xml:"len,attr"`
}

// cur hand posOffset
type WordPosOffset struct {
	XMLName xml.Name `xml:"posOffset"`
}

// cur hand stretch
type WordStretch struct {
	XMLName  xml.Name     `xml:"stretch"`
	FillRect WordFillRect `xml:"fillRect"`
}

// cur hand imagedata
type WordImagedata struct {
	XMLName xml.Name `xml:"imagedata"`
	Id      string   `xml:"id,attr"`
	Title   string   `xml:"title,attr"`
}

// cur hand p
type WordP struct {
	XMLName       xml.Name            `xml:"p"`
	PPr           WordPPr             `xml:"pPr"`
	BookmarkStart []WordBookmarkStart `xml:"bookmarkStart"`
	BookmarkEnd   []WordBookmarkEnd   `xml:"bookmarkEnd"`
	R             []WordR             `xml:"r"`
}

// cur hand txbx
type WordTxbx struct {
	XMLName     xml.Name        `xml:"txbx"`
	TxbxContent WordTxbxContent `xml:"txbxContent"`
}

// cur hand highlight
type WordHighlight struct {
	XMLName xml.Name `xml:"highlight"`
	Val     string   `xml:"val,attr"`
}

// cur hand tblInd
type WordTblInd struct {
	XMLName xml.Name `xml:"tblInd"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand tr
type WordTr struct {
	XMLName xml.Name    `xml:"tr"`
	TrPr    WordTrPr    `xml:"trPr"`
	Tc      []WordTc    `xml:"tc"`
	TblPrEx WordTblPrEx `xml:"tblPrEx"`
}

// cur hand ilvl
type WordIlvl struct {
	XMLName xml.Name `xml:"ilvl"`
	Val     string   `xml:"val,attr"`
}

// cur hand numId
type WordNumId struct {
	XMLName xml.Name `xml:"numId"`
	Val     string   `xml:"val,attr"`
}

// cur hand pict
type WordPict struct {
	XMLName xml.Name  `xml:"pict"`
	Shape   WordShape `xml:"shape"`
}

// cur hand LockedField
type WordLockedField struct {
	XMLName xml.Name `xml:"LockedField"`
}

// cur hand top
type WordTop struct {
	XMLName xml.Name `xml:"top"`
	W       string   `xml:"w,attr"`
	Type    string   `xml:"type,attr"`
}

// cur hand outlineLvl
type WordOutlineLvl struct {
	XMLName xml.Name `xml:"outlineLvl"`
	Val     string   `xml:"val,attr"`
}

// cur hand autoSpaceDN
type WordAutoSpaceDN struct {
	XMLName xml.Name `xml:"autoSpaceDN"`
	Val     string   `xml:"val,attr"`
}

// cur hand graphic
type WordGraphic struct {
	XMLName     xml.Name        `xml:"graphic"`
	GraphicData WordGraphicData `xml:"graphicData"`
	A           string          `xml:"a,attr"`
}

// cur hand tab
type WordTab struct {
	XMLName xml.Name `xml:"tab"`
	Val     string   `xml:"val,attr"`
	Pos     string   `xml:"pos,attr"`
}

// cur hand tblPrEx
type WordTblPrEx struct {
	XMLName    xml.Name       `xml:"tblPrEx"`
	TblCellMar WordTblCellMar `xml:"tblCellMar"`
	TblLayout  WordTblLayout  `xml:"tblLayout"`
}

// cur hand rFonts
type WordRFonts struct {
	XMLName xml.Name `xml:"rFonts"`
	Hint    string   `xml:"hint,attr"`
}

// cur hand picLocks
type WordPicLocks struct {
	XMLName        xml.Name `xml:"picLocks"`
	NoChangeAspect string   `xml:"noChangeAspect,attr"`
}

// cur hand blipFill
type WordBlipFill struct {
	XMLName xml.Name    `xml:"blipFill"`
	Blip    WordBlip    `xml:"blip"`
	Stretch WordStretch `xml:"stretch"`
}

// cur hand graphicData
type WordGraphicData struct {
	XMLName xml.Name `xml:"graphicData"`
	Pic     WordPic  `xml:"pic"`
	Uri     string   `xml:"uri,attr"`
}

//add w: to every line
func EncodeWord(v interface{}, prefix string) string {

	b, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func DecodeWord(filename string) *WordDocument {
	content, err := ioutil.ReadFile(filename)
	var v WordDocument
	//fmt.Printf("%s\n", string(content))
	err = xml.Unmarshal(content, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
	return &v
}
