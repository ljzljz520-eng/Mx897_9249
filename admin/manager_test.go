package admin

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"path/filepath"
	"testing"
)

func TestManagerEligibility(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	m := New(s)
	s.SaveReader(model.NewReader("R", "A", "D", 0, "idle"))
	if e := m.OpenBorrowing("R", "admin"); e != nil {
		t.Fatal(e)
	}
	ok, _ := m.IsEligible("R")
	if !ok {
		t.Fatal("not eligible")
	}
}
