package storage

import (
	"librarytransfer/model"
	"path/filepath"
	"testing"
)

func TestStoreReaderRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := model.NewReader("R1", "A", "CS", 0, "idle")
	if e = s.SaveReader(r); e != nil {
		t.Fatal(e)
	}
	got, e := s.GetReader("R1")
	if e != nil || got.Name != "A" {
		t.Fatalf("%+v %v", got, e)
	}
}
func TestStoreLogs(t *testing.T) {
	s, _ := Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	s.SaveLog(model.TaskLog{TaskID: "t", Level: "info", Message: "m"})
	ls, e := s.Logs("t")
	if e != nil || len(ls) != 1 {
		t.Fatal(e)
	}
}
