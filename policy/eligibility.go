package policy

import "librarytransfer/model"

type Decision struct {
	Allowed bool
	Reasons []string
}

func Evaluate(r model.Reader) Decision {
	reasons := []string{}
	if r.CardNumber == "" {
		reasons = append(reasons, "missing card")
	}
	if r.Name == "" {
		reasons = append(reasons, "missing name")
	}
	if r.DebtCents != 0 {
		reasons = append(reasons, "debt outstanding")
	}
	if r.BorrowingStatus != "idle" {
		reasons = append(reasons, "not idle")
	}
	return Decision{Allowed: len(reasons) == 0, Reasons: reasons}
}
func CanTransfer(r model.Reader) bool { return len(Check(r)) == 0 }
func CanEnable(r model.Reader) bool   { return Evaluate(r).Allowed }
func Reasons(r model.Reader) []string { return Evaluate(r).Reasons }
func RuleNames() []string {
	out := []string{}
	for _, r := range DefaultRules() {
		out = append(out, r.Name)
	}
	return out
}
func CountFailures(readers []model.Reader) int {
	n := 0
	for _, r := range readers {
		if !CanTransfer(r) {
			n++
		}
	}
	return n
}
