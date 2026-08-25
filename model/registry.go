package model

type ReaderRegistry struct{ items map[string]Reader }

func NewRegistry() *ReaderRegistry { return &ReaderRegistry{items: map[string]Reader{}} }
func (r *ReaderRegistry) Add(v Reader) bool {
	if v.CardNumber == "" || r.items == nil {
		return false
	}
	if _, ok := r.items[v.CardNumber]; ok {
		return false
	}
	r.items[v.CardNumber] = v
	return true
}
func (r *ReaderRegistry) Upsert(v Reader) bool {
	if v.CardNumber == "" {
		return false
	}
	if r.items == nil {
		r.items = map[string]Reader{}
	}
	_, ex := r.items[v.CardNumber]
	r.items[v.CardNumber] = v
	return !ex
}
func (r *ReaderRegistry) Get(k string) (Reader, bool) { v, ok := r.items[k]; return v, ok }
func (r *ReaderRegistry) Remove(k string) bool {
	if _, ok := r.items[k]; !ok {
		return false
	}
	delete(r.items, k)
	return true
}
func (r *ReaderRegistry) Len() int { return len(r.items) }
func (r *ReaderRegistry) Values() []Reader {
	out := []Reader{}
	for _, v := range r.items {
		out = append(out, v)
	}
	return out
}
func (r *ReaderRegistry) Clear() { r.items = map[string]Reader{} }
func (r *ReaderRegistry) Enabled() []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if v.Enabled {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) Debt() int64 {
	var n int64
	for _, v := range r.items {
		n += v.DebtCents
	}
	return n
}
func (r *ReaderRegistry) ByDepartment(d string) []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if v.Department == d {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) Borrowed() []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if v.IsBorrowed() {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) Eligible() []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if v.CanBorrow() {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) Replace(rs []Reader) {
	r.Clear()
	for _, v := range rs {
		r.Upsert(v)
	}
}
func (r *ReaderRegistry) Merge(other *ReaderRegistry) int {
	n := 0
	for _, v := range other.Values() {
		if r.Upsert(v) {
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) Cards() []string {
	out := []string{}
	for k := range r.items {
		out = append(out, k)
	}
	return out
}
func (r *ReaderRegistry) HasDebt() bool {
	for _, v := range r.items {
		if v.DebtCents > 0 {
			return true
		}
	}
	return false
}
func (r *ReaderRegistry) AllValid() bool {
	for _, v := range r.items {
		if !v.Valid() {
			return false
		}
	}
	return true
}
func (r *ReaderRegistry) Invalid() []Reader {
	out := []Reader{}
	for _, v := range r.items {
		if !v.Valid() {
			out = append(out, v)
		}
	}
	return out
}
func (r *ReaderRegistry) CountStatus(status string) int {
	n := 0
	for _, v := range r.items {
		if v.BorrowingStatus == status {
			n++
		}
	}
	return n
}
func (r *ReaderRegistry) Toggle(card string) bool {
	v, ok := r.items[card]
	if !ok {
		return false
	}
	v.Enabled = !v.Enabled
	r.items[card] = v
	return v.Enabled
}
