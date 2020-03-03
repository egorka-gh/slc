package post

import (
	"encoding/xml"
	"io"
)

//Package - root node
type Package struct {
	XMLName xml.Name `xml:"PACKAGE"`
	Name    string   `xml:"name,attr"`
	Object  *Object
}

//Encode write xml to w
func (p *Package) Encode(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(p)
}

//Object - document root node
type Object struct {
	XMLName     xml.Name `xml:"POSTOBJECT"`
	Description string   `xml:"description,attr"`
	Action      string   `xml:"action,attr"`
	ID          string   `xml:"Id"`
	Object      interface{}
}

//DocIW - document type IW node
type DocIW struct {
	XMLName  xml.Name `xml:"IW"`
	Document *Document
	WayBill  *WayBill
	DocBases []DocBase
	DocProps []DocProp
	Spec     []Spec
	SpecBY   []SpecBY
}

//Document - SMDOCUMENTS table node
type Document struct {
	XMLName           xml.Name `xml:"SMDOCUMENTS"`
	ID                string   `xml:"ID"`
	DOCTYPE           string   `xml:"DOCTYPE"`
	BORNIN            string   `xml:"BORNIN"`
	CREATEDAT         string   `xml:"CREATEDAT"`
	CURRENCYMULTORDER int      `xml:"CURRENCYMULTORDER"`
	CURRENCYRATE      float64  `xml:"CURRENCYRATE"`
	CURRENCYTYPE      int      `xml:"CURRENCYTYPE"`
	DOCSTATE          int      `xml:"DOCSTATE"`
	ISROUBLES         int      `xml:"ISROUBLES"`
	LOCATIONFROM      int      `xml:"LOCATIONFROM"`
	LOCATIONTO        int      `xml:"LOCATIONTO"`
	OPCODE            int      `xml:"OPCODE"`
	PRICEROUNDMODE    int      `xml:"PRICEROUNDMODE"`
	TOTAL             float64  `xml:"-"`
	TOTALSUM          string   `xml:"TOTALSUM"`
	TOTALSUMCUR       float64  `xml:"TOTALSUMCUR"`
}

//DocBase - SMCOMMONBASES table node
type DocBase struct {
	XMLName     xml.Name `xml:"SMCOMMONBASES"`
	ID          string   `xml:"ID"`
	DOCTYPE     string   `xml:"DOCTYPE"`
	BASEID      string   `xml:"BASEID"`
	BASEDOCTYPE string   `xml:"BASEDOCTYPE"`
}

//DocProp - SMDOCPROPS table node
type DocProp struct {
	XMLName    xml.Name `xml:"SMDOCPROPS"`
	DOCID      string   `xml:"DOCID"`
	DOCTYPE    string   `xml:"DOCTYPE"`
	PARAMNAME  string   `xml:"PARAMNAME"`
	PARAMVALUE string   `xml:"PARAMVALUE"`
}

//Spec - SMSPEC table node
type Spec struct {
	XMLName         xml.Name `xml:"SMSPEC"`
	DOCID           string   `xml:"DOCID"`
	DOCTYPE         string   `xml:"DOCTYPE"`
	SPECITEM        int      `xml:"SPECITEM"`
	ARTICLE         string   `xml:"ARTICLE"`
	DISPLAYITEM     int      `xml:"DISPLAYITEM"`
	ITEMPRICE       string   `xml:"ITEMPRICE"`
	ITEMPRICECUR    float64  `xml:"ITEMPRICECUR"`
	ITEMPRICENOTAX  string   `xml:"ITEMPRICENOTAX"`
	QUANTITY        string   `xml:"QUANTITY"`
	TOTALPRICE      string   `xml:"TOTALPRICE"`
	TOTALPRICECUR   float64  `xml:"TOTALPRICECUR"`
	TOTALPRICENOTAX string   `xml:"TOTALPRICENOTAX"`
}

//SpecBY - SMSPECBY table node
type SpecBY struct {
	XMLName            xml.Name `xml:"SMSPECBY"`
	DOCID              string   `xml:"DOCID"`
	DOCTYPE            string   `xml:"DOCTYPE"`
	SPECITEM           int      `xml:"SPECITEM"`
	STATEREGULATION    int      `xml:"STATEREGULATION"`
	MANUFACTURERSPRICE string   `xml:"MANUFACTURERSPRICE"`
}

//WayBill - SMINTERNALWAYBILLS table node
type WayBill struct {
	XMLName       xml.Name `xml:"SMINTERNALWAYBILLS"`
	ID            string   `xml:"ID"`
	DOCTYPE       string   `xml:"DOCTYPE"`
	OURSELFCLIENT int      `xml:"OURSELFCLIENT"`
}
