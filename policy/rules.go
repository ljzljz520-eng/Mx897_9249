package policy

import (
	"fmt"
	"librarytransfer/model"
)

type Rule struct {
	Name    string
	Check   func(model.Reader) bool
	Message string
}

func DefaultRules() []Rule {
	return []Rule{{"card", func(r model.Reader) bool { return r.CardNumber != "" }, "card number required"}, {"name", func(r model.Reader) bool { return r.Name != "" }, "name required"}, {"department", func(r model.Reader) bool { return r.Department != "" }, "department required"}, {"debt", func(r model.Reader) bool { return r.DebtCents >= 0 }, "debt must be nonnegative"}, {"status", func(r model.Reader) bool { return r.BorrowingStatus == "idle" || r.BorrowingStatus == "borrowed" }, "status invalid"}}
}
func Check(r model.Reader) []string {
	out := []string{}
	for _, rule := range DefaultRules() {
		if !rule.Check(r) {
			out = append(out, rule.Message)
		}
	}
	return out
}
func Explain(r model.Reader) string {
	issues := Check(r)
	if len(issues) == 0 {
		return "valid"
	}
	return fmt.Sprintf("%d policy violations", len(issues))
}
func CanPublish(r model.Reader) bool { return len(Check(r)) == 0 && r.DebtCents == 0 }
