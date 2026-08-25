package storage

import (
	"go.etcd.io/bbolt"
	"librarytransfer/model"
	"time"
)

func (s *Store) UpdateReaderAndTask(r model.Reader, t model.TransferTask) error {
	s.tasksMu.Lock()
	defer s.tasksMu.Unlock()
	s.readersMu.Lock()
	defer s.readersMu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error {
		if e := put(tx, []byte("readers"), r.CardNumber, r); e != nil {
			return e
		}
		return put(tx, []byte("tasks"), t.ID, t)
	})
}
func (s *Store) MarkEnabled(card string, enabled bool) error {
	r, e := s.GetReader(card)
	if e != nil {
		return e
	}
	r.Enabled = enabled
	r.UpdatedAt = time.Unix(0, 0).UTC()
	return s.SaveReader(r)
}
