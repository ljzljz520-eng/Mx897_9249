package codec

import (
	"encoding/json"
	"fmt"
	"librarytransfer/model"
)

func EncodeReader(r model.Reader) ([]byte, error) { return json.Marshal(r) }
func DecodeReader(b []byte) (model.Reader, error) {
	var r model.Reader
	e := json.Unmarshal(b, &r)
	return r, e
}
func EncodeTask(t model.TransferTask) ([]byte, error) { return json.Marshal(t) }
func DecodeTask(b []byte) (model.TransferTask, error) {
	var t model.TransferTask
	e := json.Unmarshal(b, &t)
	return t, e
}
func MustReader(b []byte) model.Reader {
	r, e := DecodeReader(b)
	if e != nil {
		panic(fmt.Sprintf("reader decode: %v", e))
	}
	return r
}
