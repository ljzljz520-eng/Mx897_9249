package validate

import (
	"librarytransfer/model"
	"testing"
)

func TestValidateBatchDuplicate(t *testing.T) {
	r := model.NewReader("R", "A", "D", 0, "idle")
	if len(ValidateBatch([]model.Reader{r, r})) == 0 {
		t.Fatal("expected duplicate")
	}
}
func TestParseDebt(t *testing.T) {
	v, e := model.ParseDebt("12.34")
	if e != nil || v != 1234 {
		t.Fatalf("%d %v", v, e)
	}
}
