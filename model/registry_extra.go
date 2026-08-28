package model

func (r *ReaderRegistry) Names() []string {
	out := []string{}
	for _, v := range r.items {
		out = append(out, v.Name)
	}
	return out
}
func (r *ReaderRegistry) Departments() []string {
	out := []string{}
	seen := map[string]bool{}
	for _, v := range r.items {
		if !seen[v.Department] {
			seen[v.Department] = true
			out = append(out, v.Department)
		}
	}
	return out
}
func (r *ReaderRegistry) ClearDebt() int {
	n := 0
	for k, v := range r.items {
		if v.DebtCents != 0 {
			v.DebtCents = 0
			r.items[k] = v
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) SetDepartment(from, to string) int {
	n := 0
	for k, v := range r.items {
		if v.Department == from {
			v.Department = to
			r.items[k] = v
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) SetStatus(card, status string) bool {
	v, ok := r.items[card]
	if !ok {
		return false
	}
	v.BorrowingStatus = status
	r.items[card] = v
	return true
}
func (r *ReaderRegistry) SetName(card, name string) bool {
	v, ok := r.items[card]
	if !ok {
		return false
	}
	v.Name = name
	r.items[card] = v
	return true
}
func (r *ReaderRegistry) SetDebt(card string, debt int64) bool {
	v, ok := r.items[card]
	if !ok {
		return false
	}
	v.DebtCents = debt
	r.items[card] = v
	return true
}
func (r *ReaderRegistry) Find(predicate func(Reader) bool) []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if predicate(v) {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) Count(predicate func(Reader) bool) int { return len(r.Find(predicate)) }
func (r *ReaderRegistry) Apply(fn func(Reader) Reader) {
	for k, v := range r.items {
		r.items[k] = fn(v)
	}
}
func (r *ReaderRegistry) Snapshot() map[string]Reader {
	out := map[string]Reader{}
	for k, v := range r.items {
		out[k] = v
	}
	return out
}
func (r *ReaderRegistry) Restore(snapshot map[string]Reader) {
	r.items = map[string]Reader{}
	for k, v := range snapshot {
		r.items[k] = v
	}
}
func (r *ReaderRegistry) Missing(names []string) []string {
	out := []string{}
	for _, k := range names {
		if _, ok := r.items[k]; !ok {
			out = append(out, k)
		}
	}
	return out
}
func (r *ReaderRegistry) DuplicateCards(rs []Reader) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, v := range rs {
		if seen[v.CardNumber] {
			out = append(out, v.CardNumber)
		}
		seen[v.CardNumber] = true
	}
	return out
}
func (r *ReaderRegistry) MergePreferExisting(other *ReaderRegistry) int {
	n := 0
	for k, v := range other.items {
		if _, ok := r.items[k]; !ok {
			r.items[k] = v
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) MergePreferIncoming(other *ReaderRegistry) int {
	n := 0
	for k, v := range other.items {
		if old, ok := r.items[k]; !ok || !old.Equal(v) {
			r.items[k] = v
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) CountEligible() int { return len(r.Eligible()) }
func (r *ReaderRegistry) CountBorrowed() int { return len(r.Borrowed()) }
func (r *ReaderRegistry) CountEnabled() int  { return len(r.Enabled()) }
func (r *ReaderRegistry) CountDebt() int {
	return r.Count(func(v Reader) bool { return v.DebtCents > 0 })
}
