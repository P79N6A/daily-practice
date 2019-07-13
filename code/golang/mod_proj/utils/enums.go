package utils

type UserStatus int

const (
	Normal UserStatus = iota
	Disabled = -1
)