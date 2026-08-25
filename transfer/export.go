package transfer

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/storage"
)

func ExportTask(s *storage.Store, id string) (string, error) {
	t, e := s.GetTask(id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s,%s,%d,%d,%d", t.ID, t.Status, t.Total, t.Accepted, t.Rejected), nil
}
func ExportLogs(s *storage.Store, id string) []string {
	ls, e := s.Logs(id)
	if e != nil {
		return nil
	}
	out := []string{}
	for _, l := range ls {
		out = append(out, l.Level+":"+l.Message)
	}
	return out
}
func ImportAndValidate(rows [][]string) ([]model.Reader, []string, error) {
	rs, e := ParseLegacy(rows)
	if e != nil {
		return nil, nil, e
	}
	steps := New(nil).Prepare(rs)
	notes := []string{}
	for _, x := range steps {
		notes = append(notes, x.Name+":"+x.Detail)
	}
	return rs, notes, nil
}
func NormalizeTask(t model.TransferTask) model.TransferTask {
	if t.Status == "" {
		t.Status = model.TaskPending
	}
	if t.Total < 0 {
		t.Total = 0
	}
	if t.Accepted < 0 {
		t.Accepted = 0
	}
	if t.Rejected < 0 {
		t.Rejected = 0
	}
	return t
}
