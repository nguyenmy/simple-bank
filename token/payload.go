package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrorExpiredToken = errors.New("token has expired")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issueAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewPayload(username string, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Valid() error {
	fmt.Println("valid is invoked")
	if time.Now().After(p.ExpiredAt) {
		return ErrorExpiredToken
	}

	return nil
}
