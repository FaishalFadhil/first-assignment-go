package models

import "fmt"

type Bio struct {
	ID      int
	name    string
	address string
	job     string
	reason  string
}

func (b Bio) setReason(name string, reason string) string {
	result := fmt.Sprintf("%s %s %s", name, "wants to", reason)
	return result
}
