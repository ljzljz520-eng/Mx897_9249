package importer

import (
	"encoding/csv"
	"io"
	"librarytransfer/model"
	"librarytransfer/transfer"
)

func ReadCSV(r io.Reader) ([]model.Reader, error) {
	rows, e := csv.NewReader(r).ReadAll()
	if e != nil {
		return nil, e
	}
	if len(rows) > 0 && rows[0][0] == "card_number" {
		rows = rows[1:]
	}
	return transfer.ParseLegacy(rows)
}
func WriteCSV(w io.Writer, readers []model.Reader) error {
	cw := csv.NewWriter(w)
	if e := cw.Write([]string{"card_number", "name", "department", "debt", "status"}); e != nil {
		return e
	}
	for _, row := range transfer.ExportReaders(readers) {
		if e := cw.Write(row); e != nil {
			return e
		}
	}
	cw.Flush()
	return cw.Error()
}
