package dto

type ContextKey string

var (
	ContextKeySession ContextKey = "session"
	ContextKeyToken   ContextKey = "token"
)
