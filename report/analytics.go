package report

import (
	"librarytransfer/model"
	"sort"
)

type Analytics struct {
	Total, Enabled, DebtFree, Borrowed int
	Debt                               int64
	Departments                        map[string]int
}

func Analyze(rs []model.Reader) Analytics {
	a := Analytics{Total: len(rs), Departments: map[string]int{}}
	for _, r := range rs {
		a.Departments[r.Department]++
		a.Debt += r.DebtCents
		if r.Enabled {
			a.Enabled++
		}
		if r.DebtCents == 0 {
			a.DebtFree++
		}
		if r.BorrowingStatus == "borrowed" {
			a.Borrowed++
		}
	}
	return a
}
func (a Analytics) EligibilityRate() float64 {
	if a.Total == 0 {
		return 0
	}
	return float64(a.DebtFree) / float64(a.Total)
}
func (a Analytics) EnabledRate() float64 {
	if a.Total == 0 {
		return 0
	}
	return float64(a.Enabled) / float64(a.Total)
}
func (a Analytics) DepartmentNames() []string {
	out := []string{}
	for k := range a.Departments {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
func (a Analytics) DepartmentCount(d string) int { return a.Departments[d] }
func (a Analytics) HasOverdue() bool             { return a.Debt > 0 }
func (a Analytics) Empty() bool                  { return a.Total == 0 }
func (a Analytics) Copy() Analytics {
	b := Analytics{Total: a.Total, Enabled: a.Enabled, DebtFree: a.DebtFree, Borrowed: a.Borrowed, Debt: a.Debt, Departments: map[string]int{}}
	for k, v := range a.Departments {
		b.Departments[k] = v
	}
	return b
}
func (a Analytics) Add(b Analytics) Analytics {
	c := a.Copy()
	c.Total += b.Total
	c.Enabled += b.Enabled
	c.DebtFree += b.DebtFree
	c.Borrowed += b.Borrowed
	c.Debt += b.Debt
	for k, v := range b.Departments {
		c.Departments[k] += v
	}
	return c
}
func (a Analytics) Difference(b Analytics) Analytics {
	c := a.Copy()
	c.Total -= b.Total
	c.Enabled -= b.Enabled
	c.DebtFree -= b.DebtFree
	c.Borrowed -= b.Borrowed
	c.Debt -= b.Debt
	return c
}
func (a Analytics) Valid() bool {
	return a.Total >= 0 && a.Enabled >= 0 && a.DebtFree >= 0 && a.Borrowed >= 0
}
