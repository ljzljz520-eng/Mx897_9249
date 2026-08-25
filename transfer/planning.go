package transfer

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/validate"
)

type Plan struct {
	TaskID  string
	Readers []model.Reader
	Steps   []string
	Issues  []validate.Issue
}

func NewPlan(id string, rs []model.Reader) Plan {
	return Plan{TaskID: id, Readers: append([]model.Reader{}, rs...), Steps: []string{"connect", "map", "validate", "persist", "confirm"}}
}
func (p Plan) Valid() bool { return len(p.Issues) == 0 && len(p.Readers) > 0 }
func (p *Plan) Validate() []validate.Issue {
	p.Issues = validate.ValidateBatch(p.Readers)
	return p.Issues
}
func (p Plan) StepCount() int { return len(p.Steps) }
func (p Plan) Step(i int) string {
	if i < 0 || i >= len(p.Steps) {
		return ""
	}
	return p.Steps[i]
}
func (p Plan) Description() string {
	return fmt.Sprintf("%s: %d readers across %d steps", p.TaskID, len(p.Readers), len(p.Steps))
}
func (p Plan) Cards() []string {
	out := []string{}
	for _, r := range p.Readers {
		out = append(out, r.CardNumber)
	}
	return out
}
func (p Plan) DebtTotal() int64 {
	var n int64
	for _, r := range p.Readers {
		n += r.DebtCents
	}
	return n
}
func (p Plan) HasDuplicates() bool       { return len(validate.ValidateBatch(p.Readers)) > 0 }
func (p *Plan) AddReader(r model.Reader) { p.Readers = append(p.Readers, r) }
func (p *Plan) AddStep(step string) {
	if step != "" {
		p.Steps = append(p.Steps, step)
	}
}
func (p *Plan) RemoveReader(card string) int {
	out := []model.Reader{}
	n := 0
	for _, r := range p.Readers {
		if r.CardNumber == card {
			n++
		} else {
			out = append(out, r)
		}
	}
	p.Readers = out
	return n
}
func (p Plan) ReadyForPersist() bool {
	for _, r := range p.Readers {
		if !r.Valid() {
			return false
		}
	}
	return true
}
func (p Plan) ReadyForConfirm() bool { return p.Valid() && p.ReadyForPersist() }
func (p Plan) Summary() map[string]int {
	return map[string]int{"readers": len(p.Readers), "steps": len(p.Steps), "issues": len(p.Issues)}
}
