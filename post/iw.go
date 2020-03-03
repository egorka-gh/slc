package post

import (
	"fmt"
	"strconv"
)

//helper for iw document

//NewIW creates new IW document
func NewIW(id, created string, from, to, client, opcode, state int) *DocIW {
	iw := &DocIW{
		Document: &Document{
			ID:                id,
			DOCTYPE:           "IW",
			CURRENCYMULTORDER: 0,
			CURRENCYRATE:      1,
			CURRENCYTYPE:      0,
			ISROUBLES:         1,
			PRICEROUNDMODE:    0,
			OPCODE:            opcode,
			DOCSTATE:          state,
			CREATEDAT:         created,
			LOCATIONFROM:      from,
			LOCATIONTO:        to,
		},
		WayBill: &WayBill{
			DOCTYPE:       "IW",
			ID:            id,
			OURSELFCLIENT: client,
		},
		DocBases: make([]DocBase, 0, 5),
		DocProps: make([]DocProp, 0, 10),
		Spec:     make([]Spec, 0, 50),
		SpecBY:   make([]SpecBY, 0, 50),
	}
	return iw
}

//AddBaseDoc add
func (iw *DocIW) AddBaseDoc(doctype, id string) {
	b := DocBase{
		ID:          iw.Document.ID,
		DOCTYPE:     iw.Document.DOCTYPE,
		BASEDOCTYPE: doctype,
		BASEID:      id,
	}
	iw.DocBases = append(iw.DocBases, b)
}

//AddProp add
func (iw *DocIW) AddProp(name, value string) {
	p := DocProp{
		DOCID:      iw.Document.ID,
		DOCTYPE:    iw.Document.DOCTYPE,
		PARAMNAME:  name,
		PARAMVALUE: value,
	}
	iw.DocProps = append(iw.DocProps, p)
}

//AddSpecItem add
func (iw *DocIW) AddSpecItem(article string, qtty, price float64) {
	s := Spec{
		DOCID:           iw.Document.ID,
		DOCTYPE:         iw.Document.DOCTYPE,
		SPECITEM:        len(iw.Spec) + 1,
		DISPLAYITEM:     len(iw.Spec) + 1,
		ARTICLE:         article,
		ITEMPRICE:       strconv.FormatFloat(price, 'f', 2, 64),
		ITEMPRICENOTAX:  strconv.FormatFloat(price, 'f', 2, 64),
		QUANTITY:        strconv.FormatFloat(qtty, 'f', 3, 64),
		TOTALPRICE:      strconv.FormatFloat(qtty*price, 'f', 2, 64),
		TOTALPRICENOTAX: strconv.FormatFloat(qtty*price, 'f', 2, 64),
	}
	iw.Spec = append(iw.Spec, s)
	sby := SpecBY{
		DOCID:              iw.Document.ID,
		DOCTYPE:            iw.Document.DOCTYPE,
		SPECITEM:           s.SPECITEM,
		MANUFACTURERSPRICE: strconv.FormatFloat(price, 'f', 2, 64),
	}
	iw.SpecBY = append(iw.SpecBY, sby)
	iw.Document.TOTAL += qtty * price
	iw.Document.TOTALSUM = strconv.FormatFloat(iw.Document.TOTAL, 'f', 2, 64)
}

//CreatePackage create
func (iw *DocIW) CreatePackage(born string) *Package {
	iw.Document.BORNIN = born
	o := Object{
		Action:      "normal",
		Description: "Накладная на перемещение",
		ID:          fmt.Sprintf("%s%s", "IW", iw.Document.ID),
		Object:      iw,
	}
	p := Package{
		Name:   o.ID,
		Object: &o,
	}
	return &p
}
