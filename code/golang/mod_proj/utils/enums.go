package utils

type UserStatus int

const (
	Normal UserStatus = iota
	Deleted
	Disabled = -1
)
