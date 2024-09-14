package sessions

import (
	"time"

	"github.com/Vladroon22/Test-Task-BackDev/internal/database"
)

type Session struct {
	UserID       int
	UserIP       string
	expireAt     time.Time
	RefreshToken string
	Email        string
	repo         *database.Repo
}

func NewSessions(r *database.Repo) *Session {
	return &Session{
		repo: r,
	}
}

func (s *Session) CheckSession(id int, ip string, dur time.Duration, sess *database.MySession) string {
	if sess.UserID != id {
		return "No-such-session"
	}

	if time.Now().After(sess.ExpiresAt) {
		s.DeleteSession(id)
		return "Session deleted: session expired"
	}

	if sess.UserIP != ip {
		s.DeleteSession(id)
		return "Session deleted: IP-address was changed"
	}

	return "OK"
}

func (s *Session) DeleteSession(id int) {
	s.repo.DeleteSessionFromDB(id)
}
