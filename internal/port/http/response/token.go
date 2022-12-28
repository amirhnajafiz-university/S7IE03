package response

import "time"

type TokenResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
