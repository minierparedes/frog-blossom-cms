package token

import "time"

// Maker manages tokens
// CreateToken for specific username and duration
// VerifyToken checks if token is valid
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
