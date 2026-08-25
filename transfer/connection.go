package transfer

import (
	"fmt"
	"librarytransfer/model"
	"time"
)

func CheckConnection(endpoint string) model.ConnectionCheck {
	c := model.ConnectionCheck{Endpoint: endpoint, CheckedAt: time.Unix(0, 0).UTC()}
	if endpoint == "" {
		c.Detail = "missing endpoint"
	} else if endpoint == "offline" {
		c.Detail = "unreachable"
	} else {
		c.Reachable = true
		c.Detail = "connection ready"
	}
	return c
}
func DescribeConnection(c model.ConnectionCheck) string {
	if c.Reachable {
		return fmt.Sprintf("%s: ready", c.Endpoint)
	}
	return fmt.Sprintf("%s: unavailable", c.Endpoint)
}
