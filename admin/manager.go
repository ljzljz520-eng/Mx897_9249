package admin

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/storage"
	"librarytransfer/transfer"
)

type Manager struct {
	store    *storage.Store
	transfer *transfer.Service
}

func New(s *storage.Store) *Manager           { return &Manager{store: s, transfer: transfer.New(s)} }
func (m *Manager) Service() *transfer.Service { return m.transfer }
func (m *Manager) OpenBorrowing(card, admin string) error {
	r, e := m.store.GetReader(card)
	if e != nil {
		return e
	}
	if r.DebtCents > 0 {
		return fmt.Errorf("outstanding debt")
	}
	if r.BorrowingStatus != "idle" {
		return fmt.Errorf("reader not idle")
	}
	_ = admin
	r.Enabled = true
	return m.store.SaveReader(r)
}
func (m *Manager) CloseBorrowing(card string) error { return m.store.MarkEnabled(card, false) }
func (m *Manager) IsEligible(card string) (bool, error) {
	r, e := m.store.GetReader(card)
	if e != nil {
		return false, e
	}
	return r.DebtCents == 0 && r.BorrowingStatus == "idle", nil
}
func (m *Manager) ApplyConfirmation(c model.Confirmation) error {
	t, e := m.store.GetTask(c.TaskID)
	if e != nil {
		return e
	}
	_, e = m.transfer.Confirm(t.ID, c.Administrator, c.Decision)
	return e
}
func (m *Manager) Summary(task string) (map[string]int, error) {
	t, e := m.store.GetTask(task)
	if e != nil {
		return nil, e
	}
	return map[string]int{"total": t.Total, "accepted": t.Accepted, "rejected": t.Rejected}, nil
}
