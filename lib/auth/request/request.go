package request

import (
	"bytes"
	"encoding/base64"
)

const (
	TypePassword = "password"
	TypeApiKey   = "api_key"
	TypeToken    = "token"
)

type Request interface {
	Marshal() string
	Unmarshal(string) error
}

var _ Request = (*PasswordRequest)(nil)

type PasswordRequest struct {
	Username string
	Password string
}

func (u *PasswordRequest) Marshal() string {
	return base64.URLEncoding.EncodeToString([]byte(u.Username + ":" + u.Password))
}

func (u *PasswordRequest) Unmarshal(s string) error {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	split := bytes.SplitN(data, []byte(":"), 2)
	u.Username = string(split[0])
	u.Password = string(split[1])

	return nil
}

var _ Request = (*ApiKeyRequest)(nil)

type ApiKeyRequest struct {
	ApiKey string
}

func (a *ApiKeyRequest) Marshal() string {
	return a.ApiKey
}

func (a *ApiKeyRequest) Unmarshal(s string) error {
	a.ApiKey = s
	return nil
}

type TokenRequest struct {
	Token string
}

func (t *TokenRequest) Marshal() string {
	return t.Token
}

func (t *TokenRequest) Unmarshal(s string) error {
	t.Token = s
	return nil
}
