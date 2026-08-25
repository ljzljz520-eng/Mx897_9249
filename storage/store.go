package storage

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"librarytransfer/model"
	"path/filepath"
	"sync"
)

var buckets = [][]byte{[]byte("readers"), []byte("tasks"), []byte("logs"), []byte("confirmations"), []byte("mappings")}

type Store struct {
	db                 *bbolt.DB
	readersMu, tasksMu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put(tx *bbolt.Tx, b []byte, key string, v any) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), raw)
}
func get(tx *bbolt.Tx, b []byte, key string, v any) error {
	raw := tx.Bucket(b).Get([]byte(key))
	if raw == nil {
		return fmt.Errorf("not found: %s", key)
	}
	return json.Unmarshal(raw, v)
}
func (s *Store) SaveReader(r model.Reader) error {
	s.readersMu.Lock()
	defer s.readersMu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("readers"), r.CardNumber, r) })
}
func (s *Store) GetReader(card string) (model.Reader, error) {
	s.readersMu.RLock()
	defer s.readersMu.RUnlock()
	var r model.Reader
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, []byte("readers"), card, &r) })
	return r, e
}
func (s *Store) ListReaders() ([]model.Reader, error) {
	s.readersMu.RLock()
	defer s.readersMu.RUnlock()
	out := []model.Reader{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("readers")).ForEach(func(_, v []byte) error {
			var r model.Reader
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveTask(t model.TransferTask) error {
	s.tasksMu.Lock()
	defer s.tasksMu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("tasks"), t.ID, t) })
}
func (s *Store) GetTask(id string) (model.TransferTask, error) {
	s.tasksMu.RLock()
	defer s.tasksMu.RUnlock()
	var t model.TransferTask
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, []byte("tasks"), id, &t) })
	return t, e
}
func (s *Store) SaveLog(l model.TaskLog) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("logs"), l.TaskID+fmt.Sprint(l.At.UnixNano()), l) })
}
func (s *Store) Logs(task string) ([]model.TaskLog, error) {
	out := []model.TaskLog{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("logs")).ForEach(func(_, v []byte) error {
			var l model.TaskLog
			if e := json.Unmarshal(v, &l); e != nil {
				return e
			}
			if l.TaskID == task {
				out = append(out, l)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveConfirmation(c model.Confirmation) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("confirmations"), c.TaskID, c) })
}
func (s *Store) GetConfirmation(id string) (model.Confirmation, error) {
	var c model.Confirmation
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, []byte("confirmations"), id, &c) })
	return c, e
}
func (s *Store) SaveMapping(m model.FieldMapping) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("mappings"), m.SourceField, m) })
}
func (s *Store) Mappings() ([]model.FieldMapping, error) {
	out := []model.FieldMapping{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("mappings")).ForEach(func(_, v []byte) error {
			var m model.FieldMapping
			if e := json.Unmarshal(v, &m); e != nil {
				return e
			}
			out = append(out, m)
			return nil
		})
	})
	return out, e
}
