package model

type UserRole int

const (
	UserRoleAdmin UserRole = iota
	UserRoleUser
)
