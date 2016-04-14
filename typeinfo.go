package ooxml

import "github.com/wxd237/ooxml/xml"

//cur hand rFonts
type WordRFonts struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rFonts"`
	Hint    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hint,attr"`
}

//cur hand insideV
type WordInsideV struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main insideV"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
}

//cur hand caps
type WordCaps struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main caps"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand lock
type WordLock struct {
	XMLName     xml.Name `xml:"urn:schemas-microsoft-com:office:office lock"`
	Ext         string   `xml:"urn:schemas-microsoft-com:vml ext,attr"`
	Aspectratio string   `xml:" aspectratio,attr"`
}

//cur hand cNvPr
type WordCNvPr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture cNvPr"`
	Id      string   `xml:" id,attr"`
	Name    string   `xml:" name,attr"`
	Descr   string   `xml:" descr,attr"`
}

//cur hand blip
type WordBlip struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/drawingml/2006/main blip"`
	ExtLst  WordExtLst `xml:"http://schemas.openxmlformats.org/drawingml/2006/main extLst"`
	Embed   string     `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships embed,attr"`
}

//cur hand pPr
type WordPPr struct {
	XMLName    xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
	PStyle     WordPStyle     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle"`
	SnapToGrid WordSnapToGrid `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main snapToGrid"`
	Spacing    WordSpacing    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main spacing"`
	Jc         WordJc         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main jc"`
	RPr        WordRPr        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr"`
}

//cur hand ext
type WordExt struct {
	XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/drawingml/2006/main ext"`
	UseLocalDpi WordUseLocalDpi `xml:"http://schemas.microsoft.com/office/drawing/2010/main useLocalDpi"`
	Uri         string          `xml:" uri,attr"`
}

//cur hand fillRect
type WordFillRect struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main fillRect"`
}

//cur hand cols
type WordCols struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main cols"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
}

//cur hand tcPr
type WordTcPr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcPr"`
	TcW     WordTcW  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcW"`
}

//cur hand instrText
type WordInstrText struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main instrText"`
	Space   string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
}

//cur hand shd
type WordShd struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main shd"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
	Fill    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main fill,attr"`
}

//cur hand useLocalDpi
type WordUseLocalDpi struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/drawing/2010/main useLocalDpi"`
	A14     string   `xml:"xmlns a14,attr"`
	Val     string   `xml:" val,attr"`
}

//cur hand smartTagPr
type WordSmartTagPr struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main smartTagPr"`
	Attr    []WordAttr `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main attr"`
}

//cur hand graphicData
type WordGraphicData struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphicData"`
	Pic     WordPic  `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture pic"`
	Uri     string   `xml:" uri,attr"`
}

//cur hand pStyle
type WordPStyle struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand tbl
type WordTbl struct {
	XMLName xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tbl"`
	TblPr   WordTblPr   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblPr"`
	TblGrid WordTblGrid `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblGrid"`
	Tr      []WordTr    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tr"`
}

//cur hand object
type WordObject struct {
	XMLName   xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main object"`
	Shapetype WordShapetype `xml:"urn:schemas-microsoft-com:vml shapetype"`
	Shape     WordShape     `xml:"urn:schemas-microsoft-com:vml shape"`
	OLEObject WordOLEObject `xml:"urn:schemas-microsoft-com:office:office OLEObject"`
	DxaOrig   string        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main dxaOrig,attr"`
	DyaOrig   string        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main dyaOrig,attr"`
}

//cur hand tcMar
type WordTcMar struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcMar"`
	Top     WordTop    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top"`
	Left    WordLeft   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left"`
	Bottom  WordBottom `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom"`
	Right   WordRight  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right"`
}

//cur hand pgNumType
type WordPgNumType struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgNumType"`
	Start   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main start,attr"`
}

//cur hand noWrap
type WordNoWrap struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main noWrap"`
}

//cur hand spacing
type WordSpacing struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main spacing"`
	Before      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main before,attr"`
	Line        string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main line,attr"`
	LineRule    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lineRule,attr"`
	BeforeLines string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main beforeLines,attr"`
}

//cur hand ind
type WordInd struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ind"`
	FirstLineChars string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main firstLineChars,attr"`
	FirstLine      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main firstLine,attr"`
}

//cur hand br
type WordBr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main br"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand OLEObject
type WordOLEObject struct {
	XMLName    xml.Name `xml:"urn:schemas-microsoft-com:office:office OLEObject"`
	ObjectID   string   `xml:" ObjectID,attr"`
	Id         string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Type       string   `xml:" Type,attr"`
	ProgID     string   `xml:" ProgID,attr"`
	ShapeID    string   `xml:" ShapeID,attr"`
	DrawAspect string   `xml:" DrawAspect,attr"`
}

//cur hand numId
type WordNumId struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main numId"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand vMerge
type WordVMerge struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main vMerge"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand srcRect
type WordSrcRect struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main srcRect"`
}

//cur hand szCs
type WordSzCs struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main szCs"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand tcW
type WordTcW struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcW"`
	W       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w,attr"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand iCs
type WordICs struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main iCs"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand cNvPicPr
type WordCNvPicPr struct {
	XMLName  xml.Name     `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture cNvPicPr"`
	PicLocks WordPicLocks `xml:"http://schemas.openxmlformats.org/drawingml/2006/main picLocks"`
}

//cur hand footerReference
type WordFooterReference struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main footerReference"`
	Id      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand t
type WordT struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main t"`
}

//cur hand gridCol
type WordGridCol struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main gridCol"`
	W       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w,attr"`
}

//cur hand i
type WordI struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main i"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand sectPr
type WordSectPr struct {
	XMLName         xml.Name              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sectPr"`
	HeaderReference WordHeaderReference   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main headerReference"`
	PgSz            WordPgSz              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgSz"`
	PgMar           WordPgMar             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgMar"`
	PgNumType       WordPgNumType         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgNumType"`
	Cols            WordCols              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main cols"`
	DocGrid         WordDocGrid           `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main docGrid"`
	FooterReference []WordFooterReference `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main footerReference"`
	RsidSect        string                `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidSect,attr"`
	RsidR           string                `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidR,attr"`
}

//cur hand tblW
type WordTblW struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblW"`
	W       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w,attr"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand top
type WordTop struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
}

//cur hand spPr
type WordSpPr struct {
	XMLName  xml.Name     `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture spPr"`
	Xfrm     WordXfrm     `xml:"http://schemas.openxmlformats.org/drawingml/2006/main xfrm"`
	PrstGeom WordPrstGeom `xml:"http://schemas.openxmlformats.org/drawingml/2006/main prstGeom"`
	NoFill   WordNoFill   `xml:"http://schemas.openxmlformats.org/drawingml/2006/main noFill"`
	Ln       WordLn       `xml:"http://schemas.openxmlformats.org/drawingml/2006/main ln"`
	BwMode   string       `xml:" bwMode,attr"`
}

//cur hand adjustRightInd
type WordAdjustRightInd struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main adjustRightInd"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand avLst
type WordAvLst struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main avLst"`
}

//cur hand tblBorders
type WordTblBorders struct {
	XMLName xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblBorders"`
	Right   WordRight   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right"`
	InsideH WordInsideH `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main insideH"`
	InsideV WordInsideV `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main insideV"`
	Top     WordTop     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top"`
	Left    WordLeft    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left"`
	Bottom  WordBottom  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom"`
}

//cur hand keepLines
type WordKeepLines struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main keepLines"`
}

//cur hand rStyle
type WordRStyle struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rStyle"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand off
type WordOff struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main off"`
	X       string   `xml:" x,attr"`
	Y       string   `xml:" y,attr"`
}

//cur hand ilvl
type WordIlvl struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ilvl"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand trHeight
type WordTrHeight struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main trHeight"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand tcBorders
type WordTcBorders struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcBorders"`
	Top     WordTop    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top"`
	Left    WordLeft   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left"`
	Bottom  WordBottom `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom"`
	Right   WordRight  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right"`
}

//cur hand docGrid
type WordDocGrid struct {
	XMLName   xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main docGrid"`
	Type      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
	LinePitch string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main linePitch,attr"`
}

//cur hand left
type WordLeft struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
}

//cur hand xfrm
type WordXfrm struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main xfrm"`
	Off     WordOff  `xml:"http://schemas.openxmlformats.org/drawingml/2006/main off"`
	Ext     WordExt  `xml:"http://schemas.openxmlformats.org/drawingml/2006/main ext"`
}

//cur hand inline
type WordInline struct {
	XMLName           xml.Name              `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing inline"`
	Extent            WordExtent            `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing extent"`
	EffectExtent      WordEffectExtent      `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing effectExtent"`
	DocPr             WordDocPr             `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing docPr"`
	CNvGraphicFramePr WordCNvGraphicFramePr `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing cNvGraphicFramePr"`
	Graphic           WordGraphic           `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphic"`
	DistT             string                `xml:" distT,attr"`
	DistB             string                `xml:" distB,attr"`
	DistL             string                `xml:" distL,attr"`
	DistR             string                `xml:" distR,attr"`
}

//cur hand bookmarkEnd
type WordBookmarkEnd struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkEnd"`
	Id      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main id,attr"`
}

//cur hand effectExtent
type WordEffectExtent struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing effectExtent"`
	L       string   `xml:" l,attr"`
	T       string   `xml:" t,attr"`
	R       string   `xml:" r,attr"`
	B       string   `xml:" b,attr"`
}

//cur hand vanish
type WordVanish struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main vanish"`
}

//cur hand pgSz
type WordPgSz struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgSz"`
	W       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w,attr"`
	H       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main h,attr"`
}

//cur hand stretch
type WordStretch struct {
	XMLName  xml.Name     `xml:"http://schemas.openxmlformats.org/drawingml/2006/main stretch"`
	FillRect WordFillRect `xml:"http://schemas.openxmlformats.org/drawingml/2006/main fillRect"`
}

//cur hand document
type WordDocument struct {
	XMLName   xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main document"`
	Body      WordBody `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Wpc       string   `xml:"xmlns wpc,attr"`
	M         string   `xml:"xmlns m,attr"`
	W         string   `xml:"xmlns w,attr"`
	W14       string   `xml:"xmlns w14,attr"`
	Wps       string   `xml:"xmlns wps,attr"`
	Mc        string   `xml:"xmlns mc,attr"`
	Wp        string   `xml:"xmlns wp,attr"`
	Wpg       string   `xml:"xmlns wpg,attr"`
	Ignorable string   `xml:"http://schemas.openxmlformats.org/markup-compatibility/2006 Ignorable,attr"`
	O         string   `xml:"xmlns o,attr"`
	Wp14      string   `xml:"xmlns wp14,attr"`
	Wpi       string   `xml:"xmlns wpi,attr"`
	R         string   `xml:"xmlns r,attr"`
	V         string   `xml:"xmlns v,attr"`
	W10       string   `xml:"xmlns w10,attr"`
	Wne       string   `xml:"xmlns wne,attr"`
}

//cur hand tblPr
type WordTblPr struct {
	XMLName    xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblPr"`
	TblW       WordTblW       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblW"`
	TblBorders WordTblBorders `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblBorders"`
	TblLayout  WordTblLayout  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblLayout"`
	TblLook    WordTblLook    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblLook"`
}

//cur hand f
type WordF struct {
	XMLName xml.Name `xml:"urn:schemas-microsoft-com:vml f"`
	Eqn     string   `xml:" eqn,attr"`
}

//cur hand docPr
type WordDocPr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing docPr"`
	Descr   string   `xml:" descr,attr"`
	Id      string   `xml:" id,attr"`
	Name    string   `xml:" name,attr"`
}

//cur hand picLocks
type WordPicLocks struct {
	XMLName            xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main picLocks"`
	NoChangeAspect     string   `xml:" noChangeAspect,attr"`
	NoChangeArrowheads string   `xml:" noChangeArrowheads,attr"`
}

//cur hand shapetype
type WordShapetype struct {
	XMLName        xml.Name     `xml:"urn:schemas-microsoft-com:vml shapetype"`
	Stroke         WordStroke   `xml:"urn:schemas-microsoft-com:vml stroke"`
	Formulas       WordFormulas `xml:"urn:schemas-microsoft-com:vml formulas"`
	Path           WordPath     `xml:"urn:schemas-microsoft-com:vml path"`
	Lock           WordLock     `xml:"urn:schemas-microsoft-com:office:office lock"`
	Id             string       `xml:" id,attr"`
	Coordsize      string       `xml:" coordsize,attr"`
	Spt            string       `xml:"urn:schemas-microsoft-com:office:office spt,attr"`
	Preferrelative string       `xml:"urn:schemas-microsoft-com:office:office preferrelative,attr"`
	//Path           string       `xml:" path,attr"`
	Filled  string `xml:" filled,attr"`
	Stroked string `xml:" stroked,attr"`
}

//cur hand imagedata
type WordImagedata struct {
	XMLName xml.Name `xml:"urn:schemas-microsoft-com:vml imagedata"`
	Title   string   `xml:"urn:schemas-microsoft-com:office:office title,attr"`
	Id      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
}

//cur hand drawing
type WordDrawing struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main drawing"`
	Inline  WordInline `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing inline"`
}

//cur hand rPr
type WordRPr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr"`
	Sz      WordSz   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz"`
	SzCs    WordSzCs `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main szCs"`
}

//cur hand kern
type WordKern struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main kern"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand right
type WordRight struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
}

//cur hand tc
type WordTc struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tc"`
	TcPr    WordTcPr `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tcPr"`
	P       WordP    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
}

//cur hand tblInd
type WordTblInd struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblInd"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
	W       string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w,attr"`
}

//cur hand ln
type WordLn struct {
	XMLName xml.Name   `xml:"http://schemas.openxmlformats.org/drawingml/2006/main ln"`
	NoFill  WordNoFill `xml:"http://schemas.openxmlformats.org/drawingml/2006/main noFill"`
}

//cur hand graphic
type WordGraphic struct {
	XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphic"`
	GraphicData WordGraphicData `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphicData"`
	A           string          `xml:"xmlns a,attr"`
}

//cur hand jc
type WordJc struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main jc"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand p
type WordP struct {
	XMLName       xml.Name            `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
	R             []WordR             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r"`
	BookmarkEnd   []WordBookmarkEnd   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkEnd"`
	ProofErr      []WordProofErr      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main proofErr"`
	PPr           WordPPr             `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pPr"`
	BookmarkStart []WordBookmarkStart `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkStart"`
	RsidR         string              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidR,attr"`
	RsidRDefault  string              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidRDefault,attr"`
	RsidP         string              `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidP,attr"`
}

//cur hand tabs
type WordTabs struct {
	XMLName xml.Name  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tabs"`
	Tab     []WordTab `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tab"`
}

//cur hand color
type WordColor struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand cNvGraphicFramePr
type WordCNvGraphicFramePr struct {
	XMLName           xml.Name              `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing cNvGraphicFramePr"`
	GraphicFrameLocks WordGraphicFrameLocks `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphicFrameLocks"`
}

//cur hand autoSpaceDE
type WordAutoSpaceDE struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main autoSpaceDE"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand vAlign
type WordVAlign struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main vAlign"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand snapToGrid
type WordSnapToGrid struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main snapToGrid"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand numPr
type WordNumPr struct {
	XMLName xml.Name  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main numPr"`
	Ilvl    WordIlvl  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main ilvl"`
	NumId   WordNumId `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main numId"`
}

//cur hand keepNext
type WordKeepNext struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main keepNext"`
}

//cur hand headerReference
type WordHeaderReference struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main headerReference"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
	Id      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
}

//cur hand autoSpaceDN
type WordAutoSpaceDN struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main autoSpaceDN"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand tr
type WordTr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tr"`
	Tc      []WordTc `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tc"`
	RsidR   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidR,attr"`
	RsidRPr string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidRPr,attr"`
	RsidTr  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidTr,attr"`
}

//cur hand webHidden
type WordWebHidden struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main webHidden"`
}

//cur hand hyperlink
type WordHyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink"`
	R       []WordR  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r"`
	Anchor  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main anchor,attr"`
	History string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main history,attr"`
}

//cur hand outlineLvl
type WordOutlineLvl struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main outlineLvl"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand formulas
type WordFormulas struct {
	XMLName xml.Name `xml:"urn:schemas-microsoft-com:vml formulas"`
	F       []WordF  `xml:"urn:schemas-microsoft-com:vml f"`
}

//cur hand nvPicPr
type WordNvPicPr struct {
	XMLName  xml.Name     `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture nvPicPr"`
	CNvPr    WordCNvPr    `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture cNvPr"`
	CNvPicPr WordCNvPicPr `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture cNvPicPr"`
}

//cur hand lastRenderedPageBreak
type WordLastRenderedPageBreak struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lastRenderedPageBreak"`
}

//cur hand tblLayout
type WordTblLayout struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblLayout"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand attr
type WordAttr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main attr"`
	Name    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name,attr"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand blipFill
type WordBlipFill struct {
	XMLName xml.Name    `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture blipFill"`
	Blip    WordBlip    `xml:"http://schemas.openxmlformats.org/drawingml/2006/main blip"`
	SrcRect WordSrcRect `xml:"http://schemas.openxmlformats.org/drawingml/2006/main srcRect"`
	Stretch WordStretch `xml:"http://schemas.openxmlformats.org/drawingml/2006/main stretch"`
}

//cur hand bCs
type WordBCs struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bCs"`
}

//cur hand tblGrid
type WordTblGrid struct {
	XMLName xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblGrid"`
	GridCol []WordGridCol `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main gridCol"`
}

//cur hand shape
type WordShape struct {
	XMLName   xml.Name      `xml:"urn:schemas-microsoft-com:vml shape"`
	Imagedata WordImagedata `xml:"urn:schemas-microsoft-com:vml imagedata"`
	Id        string        `xml:" id,attr"`
	Type      string        `xml:" type,attr"`
	Style     string        `xml:" style,attr"`
	Ole       string        `xml:"urn:schemas-microsoft-com:office:office ole,attr"`
}

//cur hand extent
type WordExtent struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing extent"`
	Cx      string   `xml:" cx,attr"`
	Cy      string   `xml:" cy,attr"`
}

//cur hand smartTag
type WordSmartTag struct {
	XMLName    xml.Name       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main smartTag"`
	R          WordR          `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r"`
	SmartTagPr WordSmartTagPr `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main smartTagPr"`
	Uri        string         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main uri,attr"`
	Element    string         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main element,attr"`
}

//cur hand extLst
type WordExtLst struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main extLst"`
	Ext     WordExt  `xml:"http://schemas.openxmlformats.org/drawingml/2006/main ext"`
}

//cur hand noFill
type WordNoFill struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main noFill"`
}

//cur hand pic
type WordPic struct {
	XMLName  xml.Name     `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture pic"`
	SpPr     WordSpPr     `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture spPr"`
	NvPicPr  WordNvPicPr  `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture nvPicPr"`
	BlipFill WordBlipFill `xml:"http://schemas.openxmlformats.org/drawingml/2006/picture blipFill"`
	Pic      string       `xml:"xmlns pic,attr"`
}

//cur hand r
type WordR struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r"`
	RPr     WordRPr  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr"`
	T       WordT    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main t"`
	RsidRPr string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rsidRPr,attr"`
}

//cur hand b
type WordB struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main b"`
}

//cur hand tblLook
type WordTblLook struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblLook"`
	NoVBand     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main noVBand,attr"`
	Val         string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	FirstRow    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main firstRow,attr"`
	LastRow     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lastRow,attr"`
	FirstColumn string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main firstColumn,attr"`
	LastColumn  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lastColumn,attr"`
	NoHBand     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main noHBand,attr"`
}

//cur hand stroke
type WordStroke struct {
	XMLName   xml.Name `xml:"urn:schemas-microsoft-com:vml stroke"`
	Joinstyle string   `xml:" joinstyle,attr"`
}

//cur hand path
type WordPath struct {
	XMLName         xml.Name `xml:"urn:schemas-microsoft-com:vml path"`
	Extrusionok     string   `xml:"urn:schemas-microsoft-com:office:office extrusionok,attr"`
	Gradientshapeok string   `xml:" gradientshapeok,attr"`
	Connecttype     string   `xml:"urn:schemas-microsoft-com:office:office connecttype,attr"`
}

//cur hand graphicFrameLocks
type WordGraphicFrameLocks struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphicFrameLocks"`
	NoChangeAspect string   `xml:" noChangeAspect,attr"`
	A              string   `xml:"xmlns a,attr"`
}

//cur hand lang
type WordLang struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main lang"`
	EastAsia string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main eastAsia,attr"`
}

//cur hand proofErr
type WordProofErr struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main proofErr"`
	Type    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr"`
}

//cur hand cantSplit
type WordCantSplit struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main cantSplit"`
}

//cur hand trPr
type WordTrPr struct {
	XMLName   xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main trPr"`
	CantSplit WordCantSplit `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main cantSplit"`
}

//cur hand widowControl
type WordWidowControl struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main widowControl"`
}

//cur hand sz
type WordSz struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand fldChar
type WordFldChar struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main fldChar"`
	FldCharType string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main fldCharType,attr"`
}

//cur hand prstGeom
type WordPrstGeom struct {
	XMLName xml.Name  `xml:"http://schemas.openxmlformats.org/drawingml/2006/main prstGeom"`
	AvLst   WordAvLst `xml:"http://schemas.openxmlformats.org/drawingml/2006/main avLst"`
	Prst    string    `xml:" prst,attr"`
}

//cur hand noProof
type WordNoProof struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main noProof"`
}

//cur hand tab
type WordTab struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tab"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Pos     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pos,attr"`
}

//cur hand smallCaps
type WordSmallCaps struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main smallCaps"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
}

//cur hand pgMar
type WordPgMar struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pgMar"`
	Header  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main header,attr"`
	Footer  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main footer,attr"`
	Gutter  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main gutter,attr"`
	Top     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main top,attr"`
	Right   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right,attr"`
	Bottom  string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom,attr"`
	Left    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left,attr"`
}

//cur hand bottom
type WordBottom struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bottom"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
}

//cur hand pageBreakBefore
type WordPageBreakBefore struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pageBreakBefore"`
}

//cur hand tblCellMar
type WordTblCellMar struct {
	XMLName xml.Name  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tblCellMar"`
	Left    WordLeft  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main left"`
	Right   WordRight `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main right"`
}

//cur hand bookmarkStart
type WordBookmarkStart struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkStart"`
	Id      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main id,attr"`
	Name    string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name,attr"`
}

//cur hand insideH
type WordInsideH struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main insideH"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val,attr"`
	Sz      string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,attr"`
	Space   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main space,attr"`
	Color   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,attr"`
}

//cur hand body
type WordBody struct {
	XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Tbl         []WordTbl       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main tbl"`
	BookmarkEnd WordBookmarkEnd `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkEnd"`
	SectPr      WordSectPr      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sectPr"`
	P           []WordP         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
}
