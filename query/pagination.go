package query

import "librarytransfer/model"

type Page struct {
	Items                []model.Reader
	Offset, Limit, Total int
	HasNext              bool
}

func Paginate(readers []model.Reader, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	if offset > len(readers) {
		offset = len(readers)
	}
	end := offset + limit
	if end > len(readers) {
		end = len(readers)
	}
	return Page{Items: readers[offset:end], Offset: offset, Limit: limit, Total: len(readers), HasNext: end < len(readers)}
}
func GroupByDepartment(readers []model.Reader) map[string][]model.Reader {
	out := map[string][]model.Reader{}
	for _, r := range readers {
		out[r.Department] = append(out[r.Department], r)
	}
	return out
}
func GroupByEnabled(readers []model.Reader) map[bool]int {
	out := map[bool]int{}
	for _, r := range readers {
		out[r.Enabled]++
	}
	return out
}
func Cards(readers []model.Reader) []string {
	out := []string{}
	for _, r := range readers {
		out = append(out, r.CardNumber)
	}
	return out
}
func Names(readers []model.Reader) []string {
	out := []string{}
	for _, r := range readers {
		out = append(out, r.Name)
	}
	return out
}
func FilterDebt(readers []model.Reader, min, max int64) []model.Reader {
	out := []model.Reader{}
	for _, r := range readers {
		if r.DebtCents >= min && (max < 0 || r.DebtCents <= max) {
			out = append(out, r)
		}
	}
	return out
}
