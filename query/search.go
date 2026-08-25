package query

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"strings"
)

func Search(s *storage.Store, term string) ([]model.Reader, error) {
	all, e := s.ListReaders()
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(strings.TrimSpace(term))
	out := []model.Reader{}
	for _, r := range all {
		if strings.Contains(strings.ToLower(r.CardNumber), term) || strings.Contains(strings.ToLower(r.Name), term) || strings.Contains(strings.ToLower(r.Department), term) {
			out = append(out, r)
		}
	}
	return out, nil
}
func ExactCard(s *storage.Store, card string) (model.Reader, bool) {
	r, e := s.GetReader(card)
	return r, e == nil
}
func DepartmentNames(s *storage.Store) []string {
	rs, e := s.ListReaders()
	if e != nil {
		return nil
	}
	return Names(uniqueDepartments(rs))
}
func uniqueDepartments(rs []model.Reader) []model.Reader {
	seen := map[string]bool{}
	out := []model.Reader{}
	for _, r := range rs {
		if !seen[r.Department] {
			seen[r.Department] = true
			out = append(out, r)
		}
	}
	return out
}
func EnabledReaders(s *storage.Store) []model.Reader {
	rs, e := s.ListReaders()
	if e != nil {
		return nil
	}
	out := []model.Reader{}
	for _, r := range rs {
		if r.Enabled {
			out = append(out, r)
		}
	}
	return out
}
func BorrowedReaders(s *storage.Store) []model.Reader {
	rs, e := s.ListReaders()
	if e != nil {
		return nil
	}
	out := []model.Reader{}
	for _, r := range rs {
		if r.IsBorrowed() {
			out = append(out, r)
		}
	}
	return out
}
