package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"librarytransfer/model"
)

type Inspection struct{ Readers, Tasks, Logs, Mappings int }

func (s *Store) Inspect() Inspection {
	return Inspection{Readers: s.CountReaders(), Tasks: s.countBucket("tasks"), Logs: s.countBucket("logs"), Mappings: s.countBucket("mappings")}
}
func (s *Store) countBucket(name string) int {
	n := 0
	_ = s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(name))
		if b != nil {
			n = b.Stats().KeyN
		}
		return nil
	})
	return n
}
func (i Inspection) Total() int    { return i.Readers + i.Tasks + i.Logs + i.Mappings }
func (i Inspection) Healthy() bool { return i.Readers >= 0 && i.Tasks >= 0 }
func (i Inspection) String() string {
	return fmt.Sprintf("readers=%d tasks=%d logs=%d mappings=%d", i.Readers, i.Tasks, i.Logs, i.Mappings)
}
func CloneReader(r model.Reader) model.Reader { return r }
