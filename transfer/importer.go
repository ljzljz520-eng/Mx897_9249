package transfer

import (
	"fmt"
	"librarytransfer/model"
)

func ParseLegacy(rows [][]string) ([]model.Reader, error) {
	out := []model.Reader{}
	for i, row := range rows {
		if len(row) < 5 {
			return nil, fmt.Errorf("row %d has %d fields", i, len(row))
		}
		d, e := model.ParseDebt(row[3])
		if e != nil {
			return nil, e
		}
		out = append(out, model.NewReader(row[0], row[1], row[2], d, row[4]))
	}
	return out, nil
}
func ExportReaders(readers []model.Reader) [][]string {
	out := [][]string{}
	for _, r := range readers {
		out = append(out, []string{r.CardNumber, r.Name, r.Department, model.FormatDebt(r.DebtCents), r.BorrowingStatus})
	}
	return out
}
