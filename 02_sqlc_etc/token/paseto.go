package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetMaker struct {
	paseto       *paseto.V2
	symmetrickey []byte
}

func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid KeySize")
	}

	maker := &PasetMaker{
		paseto:       paseto.NewV2(),
		symmetrickey: []byte(symmetrickey),
	}
	return maker, nil
}

func (p *PasetMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.symmetrickey, payload, nil)
}

func (p *PasetMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetrickey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
