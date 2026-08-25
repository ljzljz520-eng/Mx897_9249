package report

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/query"
	"sort"
	"strings"
)

type Report struct {
	Title     string
	Rows      []model.Reader
	TotalDebt int64
	ByStatus  map[string]int
}

func Build(title string, rows []model.Reader) Report {
	return Report{Title: title, Rows: rows, TotalDebt: query.DebtTotal(rows), ByStatus: query.StatusCounts(rows)}
}
func (r Report) CSV() string {
	var b strings.Builder
	b.WriteString("card_number,name,department,debt,status,enabled\n")
	rows := append([]model.Reader{}, r.Rows...)
	sort.Slice(rows, func(i, j int) bool { return rows[i].CardNumber < rows[j].CardNumber })
	for _, x := range rows {
		fmt.Fprintf(&b, "%s,%s,%s,%s,%s,%t\n", x.CardNumber, x.Name, x.Department, model.FormatDebt(x.DebtCents), x.BorrowingStatus, x.Enabled)
	}
	return b.String()
}
func (r Report) Summary() string {
	return fmt.Sprintf("%s: %d readers, debt %s", r.Title, len(r.Rows), model.FormatDebt(r.TotalDebt))
}
func (r Report) HasDebt() bool { return r.TotalDebt > 0 }
