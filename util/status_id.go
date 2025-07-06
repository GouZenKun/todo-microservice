package util

import (
	"strings"
)

func GetStatusID(status string) uint64 {
	s := strings.ToLower(status)
	if s == "completed" {
		return 1
	}
	if s == "pending" {
		return 2
	}
	if s == "incomplete" {
		return 3
	}
	return 0
}

// TODO : write enum
