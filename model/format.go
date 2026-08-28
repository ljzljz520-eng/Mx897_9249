package model

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatDebt(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	return fmt.Sprintf("%s%d.%02d", sign, cents/100, cents%100)
}
func ParseDebt(s string) (int64, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 2 || len(parts[1]) != 2 {
		return 0, fmt.Errorf("invalid debt %q", s)
	}
	d, e := strconv.ParseInt(parts[0], 10, 64)
	if e != nil || d < 0 {
		return 0, fmt.Errorf("invalid debt %q", s)
	}
	c, e := strconv.ParseInt(parts[1], 10, 64)
	if e != nil || c < 0 || c > 99 {
		return 0, fmt.Errorf("invalid debt %q", s)
	}
	return d*100 + c, nil
}
