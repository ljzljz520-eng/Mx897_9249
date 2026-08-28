package transfer

import (
	"fmt"
	"librarytransfer/model"
)

func (s *Service) RequireApproval(id string) error {
	t, e := s.Store.GetTask(id)
	if e != nil {
		return e
	}
	if t.Status != model.TaskAwaiting {
		return fmt.Errorf("approval not required")
	}
	return nil
}
func (s *Service) RejectIfDebt(id string, readers []model.Reader) error {
	for _, r := range readers {
		if r.DebtCents > 0 {
			_, e := s.Confirm(id, "policy", model.DecisionReject)
			return e
		}
	}
	return nil
}
func (s *Service) EnableApproved(id string, cards []string) error {
	t, e := s.Store.GetTask(id)
	if e != nil {
		return e
	}
	if !t.IsApproved() {
		return fmt.Errorf("task not approved")
	}
	for _, card := range cards {
		if e := s.Store.MarkEnabled(card, true); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) DisableAll(cards []string) error {
	for _, card := range cards {
		if e := s.Store.MarkEnabled(card, false); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) ValidateTransition(from, to string) error {
	if !model.AllowedTransition(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	return nil
}
