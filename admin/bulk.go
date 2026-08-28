package admin

import (
	"fmt"
	"librarytransfer/model"
)

func (m *Manager) OpenMany(cards []string, admin string) map[string]error {
	out := map[string]error{}
	for _, c := range cards {
		out[c] = m.OpenBorrowing(c, admin)
	}
	return out
}
func (m *Manager) CloseMany(cards []string) map[string]error {
	out := map[string]error{}
	for _, c := range cards {
		out[c] = m.CloseBorrowing(c)
	}
	return out
}
func (m *Manager) EligibleCards(cards []string) []string {
	out := []string{}
	for _, c := range cards {
		ok, e := m.IsEligible(c)
		if e == nil && ok {
			out = append(out, c)
		}
	}
	return out
}
func (m *Manager) ValidateAdministrator(name string) error {
	if len(name) < 2 {
		return fmt.Errorf("administrator name required")
	}
	return nil
}
func (m *Manager) ApprovalSummary(c model.Confirmation) string {
	if c.Approved() {
		return "approved by " + c.Administrator
	}
	return "rejected by " + c.Administrator
}
func (m *Manager) ReaderState(card string) (string, error) {
	r, e := m.store.GetReader(card)
	if e != nil {
		return "", e
	}
	return r.StateLabel(), nil
}
func (m *Manager) Debt(card string) (int64, error) {
	r, e := m.store.GetReader(card)
	if e != nil {
		return 0, e
	}
	return r.DebtCents, nil
}
func (m *Manager) SetDebt(card string, cents int64) error {
	r, e := m.store.GetReader(card)
	if e != nil {
		return e
	}
	r.DebtCents = cents
	return m.store.SaveReader(r)
}
func (m *Manager) SetDepartment(card, dept string) error {
	r, e := m.store.GetReader(card)
	if e != nil {
		return e
	}
	r.Department = dept
	return m.store.SaveReader(r)
}
func (m *Manager) SetStatus(card, status string) error {
	r, e := m.store.GetReader(card)
	if e != nil {
		return e
	}
	if status != "idle" && status != "borrowed" {
		return fmt.Errorf("invalid status")
	}
	r.BorrowingStatus = status
	return m.store.SaveReader(r)
}
