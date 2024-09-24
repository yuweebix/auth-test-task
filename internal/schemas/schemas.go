package schemas

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `db:"id"`
	Email string    `db:"email"`
}

type RefreshToken struct {
	ID     uuid.UUID `db:"id"`
	Token  string    `db:"token"`
	UserID uuid.UUID `db:"user_id"`
	IP     string    `db:"ip"`
	Expiry time.Time `db:"expires_at"`
}
