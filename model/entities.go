package model

import "time"

type Reader struct {
	CardNumber, Name, Department string
	DebtCents                    int64
	BorrowingStatus              string
	Enabled                      bool
	UpdatedAt                    time.Time
}
type TransferTask struct {
	ID, SourceName, TargetName string
	Status                     string
	CreatedAt, CompletedAt     time.Time
	Total, Accepted, Rejected  int
}
type FieldMapping struct {
	SourceField, TargetField string
	Required                 bool
}
type TaskLog struct {
	TaskID, Level, Message string
	At                     time.Time
}
type ConnectionCheck struct {
	Endpoint  string
	Reachable bool
	Detail    string
	CheckedAt time.Time
}
type Confirmation struct {
	TaskID, Administrator, Decision string
	At                              time.Time
}

func (r Reader) Valid() bool {
	return r.CardNumber != "" && r.Name != "" && r.Department != "" && r.DebtCents >= 0 && (r.BorrowingStatus == "idle" || r.BorrowingStatus == "borrowed")
}
func (r Reader) CanBorrow() bool { return r.Enabled && r.BorrowingStatus == "idle" }
func NewReader(card, name, dept string, debt int64, status string) Reader {
	return Reader{CardNumber: card, Name: name, Department: dept, DebtCents: debt, BorrowingStatus: status, UpdatedAt: time.Unix(0, 0).UTC()}
}
