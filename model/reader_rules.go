package model

import "strings"

func (r Reader) Key() string           { return strings.TrimSpace(r.CardNumber) }
func (r Reader) DepartmentKey() string { return strings.ToUpper(strings.TrimSpace(r.Department)) }
func (r Reader) DisplayName() string   { return strings.TrimSpace(r.Name) + " (" + r.Key() + ")" }
func (r Reader) DebtLabel() string {
	if r.DebtCents == 0 {
		return "clear"
	}
	return FormatDebt(r.DebtCents)
}
func (r Reader) StateLabel() string {
	if r.Enabled {
		return "open"
	}
	return "held"
}
func (r Reader) NeedsAttention() bool           { return r.DebtCents > 0 || r.BorrowingStatus == "borrowed" }
func (r Reader) IsBorrowed() bool               { return r.BorrowingStatus == "borrowed" }
func (r Reader) IsIdle() bool                   { return r.BorrowingStatus == "idle" }
func (r Reader) WithDepartment(d string) Reader { r.Department = strings.TrimSpace(d); return r }
func (r Reader) WithDebt(c int64) Reader        { r.DebtCents = c; return r }
func (r Reader) WithStatus(s string) Reader     { r.BorrowingStatus = s; return r }
func (r Reader) Enable() Reader                 { r.Enabled = true; return r }
func (r Reader) Disable() Reader                { r.Enabled = false; return r }
func (r Reader) Equal(other Reader) bool {
	return r.CardNumber == other.CardNumber && r.Name == other.Name && r.Department == other.Department && r.DebtCents == other.DebtCents && r.BorrowingStatus == other.BorrowingStatus && r.Enabled == other.Enabled
}
func (t TransferTask) IsApproved() bool { return t.Status == TaskCompleted }
func (t TransferTask) IsRejected() bool { return t.Status == TaskFailed }
func (t TransferTask) CompletionRate() float64 {
	if t.Total == 0 {
		return 0
	}
	return float64(t.Accepted) / float64(t.Total)
}
func (t TransferTask) PendingCount() int { return t.Total - t.Accepted - t.Rejected }
func (t TransferTask) CanConfirm() bool  { return t.Status == TaskAwaiting }
func (t TransferTask) Label() string     { return strings.ReplaceAll(t.Status, "_", " ") }
func (c Confirmation) Approved() bool    { return c.Decision == DecisionApprove }
func (c Confirmation) Rejected() bool    { return c.Decision == DecisionReject }
func (m FieldMapping) Key() string       { return m.SourceField + "->" + m.TargetField }
func (m FieldMapping) Valid() bool       { return m.SourceField != "" && m.TargetField != "" }
func (l TaskLog) IsError() bool          { return l.Level == "error" }
func (l TaskLog) IsInfo() bool           { return l.Level == "info" }
