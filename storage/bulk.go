package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"librarytransfer/model"
)

func (s *Store) SaveReaders(readers []model.Reader) error {
	for _, r := range readers {
		if e := s.SaveReader(r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) DeleteReader(card string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("readers")).Delete([]byte(card)) })
}

func (s *Store) SaveMappings(ms []model.FieldMapping) error {
	for _, m := range ms {
		if !m.Valid() {
			return fmt.Errorf("invalid mapping")
		}
		if e := s.SaveMapping(m); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) TaskExists(id string) bool     { _, e := s.GetTask(id); return e == nil }
func (s *Store) ReaderExists(card string) bool { _, e := s.GetReader(card); return e == nil }
func (s *Store) CountReaders() int {
	rs, e := s.ListReaders()
	if e != nil {
		return 0
	}
	return len(rs)
}
func (s *Store) CountLogs(task string) int {
	ls, e := s.Logs(task)
	if e != nil {
		return 0
	}
	return len(ls)
}
func (s *Store) AddInfo(task, msg string) error {
	return s.SaveLog(model.TaskLog{TaskID: task, Level: "info", Message: msg})
}
func (s *Store) AddError(task, msg string) error {
	return s.SaveLog(model.TaskLog{TaskID: task, Level: "error", Message: msg})
}
