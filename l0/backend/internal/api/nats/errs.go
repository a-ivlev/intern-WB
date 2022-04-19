package nats

import "errors"

var (
	ErrConnNats = errors.New("error in connecting to nats")
	ErrIncorDataNats = errors.New("incorrect data in NATS-stream, cannot parse order, drop message")
)