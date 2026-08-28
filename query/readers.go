package query

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"sort"
	"strings"
)

type ReaderQuery struct {
	Department, Status string
	EnabledOnly        bool
}

func List(s *storage.Store, q ReaderQuery) ([]model.Reader, error) {
	all, e := s.ListReaders()
	if e != nil {
		return nil, e
	}
	out := []model.Reader{}
	for _, r := range all {
		if q.Department != "" && !strings.EqualFold(q.Department, r.Department) {
			continue
		}
		if q.Status != "" && q.Status != r.BorrowingStatus {
			continue
		}
		if q.EnabledOnly && !r.Enabled {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CardNumber < out[j].CardNumber })
	return out, nil
}
func FindByName(s *storage.Store, name string) ([]model.Reader, error) { return List(s, ReaderQuery{}) }
func DebtTotal(readers []model.Reader) int64 {
	var n int64
	for _, r := range readers {
		n += r.DebtCents
	}
	return n
}
func StatusCounts(readers []model.Reader) map[string]int {
	out := map[string]int{}
	for _, r := range readers {
		out[r.BorrowingStatus]++
	}
	return out
}
