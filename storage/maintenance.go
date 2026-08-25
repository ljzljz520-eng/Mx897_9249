package storage

import (
	"fmt"
	"librarytransfer/model"
)

func (s *Store) EnsureReader(r model.Reader) error {
	if r.CardNumber == "" {
		return fmt.Errorf("card required")
	}
	if s.ReaderExists(r.CardNumber) {
		return fmt.Errorf("reader exists")
	}
	return s.SaveReader(r)
}
func (s *Store) EnsureTask(t model.TransferTask) error {
	if t.ID == "" {
		return fmt.Errorf("task id required")
	}
	if s.TaskExists(t.ID) {
		return fmt.Errorf("task exists")
	}
	return s.SaveTask(t)
}
func (s *Store) UpdateTaskStatus(id, status string) error {
	t, e := s.GetTask(id)
	if e != nil {
		return e
	}
	if !model.AllowedTransition(t.Status, status) {
		return fmt.Errorf("invalid status")
	}
	t.Status = status
	return s.SaveTask(t)
}
func (s *Store) UpdateTaskCounts(id string, accepted, rejected int) error {
	t, e := s.GetTask(id)
	if e != nil {
		return e
	}
	if accepted < 0 || rejected < 0 || accepted+rejected > t.Total {
		return fmt.Errorf("invalid counts")
	}
	t.Accepted = accepted
	t.Rejected = rejected
	return s.SaveTask(t)
}
func (s *Store) ReaderCards() ([]string, error) {
	rs, e := s.ListReaders()
	if e != nil {
		return nil, e
	}
	out := []string{}
	for _, r := range rs {
		out = append(out, r.CardNumber)
	}
	return out, nil
}
func (s *Store) DepartmentTotal(dept string) int {
	rs, e := s.ListReaders()
	if e != nil {
		return 0
	}
	n := 0
	for _, r := range rs {
		if r.Department == dept {
			n++
		}
	}
	return n
}
func (s *Store) EnabledTotal() int {
	rs, e := s.ListReaders()
	if e != nil {
		return 0
	}
	n := 0
	for _, r := range rs {
		if r.Enabled {
			n++
		}
	}
	return n
}
func (s *Store) DebtTotal() int64 {
	rs, e := s.ListReaders()
	if e != nil {
		return 0
	}
	var n int64
	for _, r := range rs {
		n += r.DebtCents
	}
	return n
}
func (s *Store) CloseAndReopen(path string) (*Store, error) {
	if e := s.Close(); e != nil {
		return nil, e
	}
	return Open(path)
}
