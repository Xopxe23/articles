package domain

import "time"

type RefreshSession struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}
