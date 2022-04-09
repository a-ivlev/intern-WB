package router

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrConnNats = errors.New("error in connecting to nats")
	ErrIncorDataNats = errors.New("incorrect data in NATS-stream, cannot parse order, drop message")
)