package report

import (
	"fmt"
	"librarytransfer/model"
)

func Headers() []string {
	return []string{"card_number", "name", "department", "debt", "status", "enabled"}
}
func Row(r model.Reader) []string {
	return []string{r.CardNumber, r.Name, r.Department, model.FormatDebt(r.DebtCents), r.BorrowingStatus, fmt.Sprint(r.Enabled)}
}
func Rows(readers []model.Reader) [][]string {
	out := [][]string{}
	for _, r := range readers {
		out = append(out, Row(r))
	}
	return out
}
func EnabledCount(readers []model.Reader) int {
	n := 0
	for _, r := range readers {
		if r.Enabled {
			n++
		}
	}
	return n
}
func BorrowedCount(readers []model.Reader) int {
	n := 0
	for _, r := range readers {
		if r.BorrowingStatus == "borrowed" {
			n++
		}
	}
	return n
}
func DebtFreeCount(readers []model.Reader) int {
	n := 0
	for _, r := range readers {
		if r.DebtCents == 0 {
			n++
		}
	}
	return n
}
