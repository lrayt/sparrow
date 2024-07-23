package runtime

import (
	"context"
	"github.com/lrayt/sparrow/core/abstract"
)

const (
	UserId string = "User-Id"
)

type Context struct {
	CTX    context.Context
	Logger abstract.Logger
}

func (c Context) Log() abstract.Logger {
	return c.Logger
}

func (c Context) UserId() string {
	if id, ok := c.CTX.Value(UserId).(string); ok {
		return id
	} else {
		return ""
	}
}
