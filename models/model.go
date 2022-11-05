package models

import "fmt"

type Bio struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func (b Bio) SetReason(name string, reason string) string {
	result := fmt.Sprintf("%s %s %s", name, "wants to", reason)
	return result
}
