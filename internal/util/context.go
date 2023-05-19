package util

import (
	"context"
)

const (
	CKUser = "user"
)

type User struct {
	ID string
	IP string
}

func Identity(c context.Context) *User {
	if u, ok := c.Value(CKUser).(*User); ok {
		return u
	}

	return &User{}
}
