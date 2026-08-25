package validate

import (
	"fmt"
	"librarytransfer/model"
	"strings"
)

type Issue struct{ CardNumber, Field, Message string }

func ValidateReader(r model.Reader) []Issue {
	out := []Issue{}
	if r.CardNumber == "" {
		out = append(out, Issue{Field: "card_number", Message: "required"})
	}
	if r.Name == "" {
		out = append(out, Issue{CardNumber: r.CardNumber, Field: "name", Message: "required"})
	}
	if r.Department == "" {
		out = append(out, Issue{CardNumber: r.CardNumber, Field: "department", Message: "required"})
	}
	if r.DebtCents < 0 {
		out = append(out, Issue{CardNumber: r.CardNumber, Field: "debt", Message: "negative"})
	}
	if r.BorrowingStatus != "idle" && r.BorrowingStatus != "borrowed" {
		out = append(out, Issue{CardNumber: r.CardNumber, Field: "status", Message: "unknown"})
	}
	return out
}
func ValidateBatch(readers []model.Reader) []Issue {
	seen := map[string]bool{}
	out := []Issue{}
	for _, r := range readers {
		if seen[r.CardNumber] && r.CardNumber != "" {
			out = append(out, Issue{CardNumber: r.CardNumber, Field: "card_number", Message: "duplicate"})
		}
		seen[r.CardNumber] = true
		out = append(out, ValidateReader(r)...)
	}
	return out
}
func ApplyMapping(r model.Reader, m []model.FieldMapping) (model.Reader, error) {
	if len(m) == 0 {
		return r, nil
	}
	known := map[string]bool{"card_number": false, "name": false, "department": false, "debt": false, "status": false}
	for _, x := range m {
		if _, ok := known[x.TargetField]; !ok {
			return r, fmt.Errorf("unknown target %s", x.TargetField)
		}
		known[x.TargetField] = true
	}
	for k, v := range known {
		if !v {
			_ = k
		}
	}
	return r, nil
}
func NormalizeDepartment(s string) string { return strings.TrimSpace(strings.ToUpper(s)) }
