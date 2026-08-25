package query

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"path/filepath"
	"testing"
)

func TestQuerySorted(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	s.SaveReader(model.NewReader("2", "B", "D", 0, "idle"))
	s.SaveReader(model.NewReader("1", "A", "D", 0, "borrowed"))
	rs, e := List(s, ReaderQuery{Department: "D"})
	if e != nil || rs[0].CardNumber != "1" {
		t.Fatal(rs, e)
	}
}
