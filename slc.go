package slc

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/egorka-gh/sm/slc/post"
)

const idLenth = 13

//Client represents slc client
type Client struct {
	//prefix to add to original id
	IDprefix string
	//SM BORNIN (database id)
	BornIn string
	//SM location from id
	LocFrom int
	//SM ourself client id
	Client int
	//SM opcode for IW document
	IWopcode int
	//SM state for IW document
	IWstate int
}

//ParseIW read and parse csv file (dispatch_result_*.csv) to Package with IW doc
func (c *Client) ParseIW(in io.Reader) (*post.Package, error) {
	r := csv.NewReader(in)
	r.Comma = '|'
	r.FieldsPerRecord = -1

	//get first line with doc header
	record, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("ParseIW: %w", err)
	}
	if len(record) < 10 {
		return nil, fmt.Errorf("ParseIW (header): wrong columns count expexted >= 10, got %d", len(record))
	}
	//build id
	id := record[6]
	//cut slc id if too long
	if len(id) > idLenth-len(c.IDprefix) {
		id = id[len(id)+len(c.IDprefix)-idLenth:]
	}
	id = c.IDprefix + id
	//doc date
	created := record[4]
	//yyyymmdd to yyyy-mm-ddT00:00:00 format
	created = created[0:4] + "-" + created[4:6] + "-" + created[6:] + "T00:00:00"
	//loc to
	to, err := strconv.Atoi(record[9])
	if err != nil {
		return nil, fmt.Errorf("ParseIW (header): location '%s' is not int", record[9])
	}
	iw := post.NewIW(id, created, c.LocFrom, to, c.Client, c.IWopcode, c.IWstate)
	iw.AddBaseDoc("SO", record[0])
	iw.AddProp("CustomLabels.User.PaperTTNNumber", record[3])

	//parse doc body
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("ParseIW (body): %w", err)
		}
		if len(record) < 8 {
			return nil, fmt.Errorf("ParseIW (body): wrong columns count expexted >= 8, got %d", len(record))
		}
		qtty, err := strconv.ParseFloat(strings.Replace(record[5], ",", ".", -1), 64)
		if err != nil {
			return nil, fmt.Errorf("ParseIW (body): quantity '%s' is not float", record[5])
		}
		price, err := strconv.ParseFloat(strings.Replace(record[7], ",", ".", -1), 64)
		if err != nil {
			return nil, fmt.Errorf("ParseIW (body): quantity '%s' is not float", record[7])
		}

		iw.AddSpecItem(record[0], qtty, price)
	}

	return iw.CreatePackage(c.BornIn), nil

}
