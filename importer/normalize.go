package importer

import (
	"librarytransfer/model"
	"strings"
)

func Normalize(r model.Reader) model.Reader {
	r.CardNumber = strings.TrimSpace(r.CardNumber)
	r.Name = strings.TrimSpace(r.Name)
	r.Department = strings.TrimSpace(r.Department)
	r.BorrowingStatus = strings.ToLower(strings.TrimSpace(r.BorrowingStatus))
	return r
}
func NormalizeAll(rs []model.Reader) []model.Reader {
	out := make([]model.Reader, 0, len(rs))
	for _, r := range rs {
		out = append(out, Normalize(r))
	}
	return out
}
func RejectBlank(rs []model.Reader) []model.Reader {
	out := []model.Reader{}
	for _, r := range rs {
		if r.CardNumber != "" {
			out = append(out, r)
		}
	}
	return out
}
func Deduplicate(rs []model.Reader) []model.Reader {
	seen := map[string]bool{}
	out := []model.Reader{}
	for _, r := range rs {
		if !seen[r.CardNumber] {
			seen[r.CardNumber] = true
			out = append(out, r)
		}
	}
	return out
}
