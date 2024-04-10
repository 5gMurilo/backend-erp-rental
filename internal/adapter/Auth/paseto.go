package Auth

import (
	"aidanwoods.dev/go-paseto"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"github.com/google/uuid"
	"strings"
	"time"
)

type PasetoToken struct {
	token    *paseto.Token
	key      *paseto.V4SymmetricKey
	parser   *paseto.Parser
	duration time.Duration
}

func New() (ports.TokenService, error) {
	duration, err := time.ParseDuration("8770h30m30s")
	if err != nil {
		return nil, err
	}

	token := paseto.NewToken()
	parser := paseto.NewParser()
	key, err := paseto.V4SymmetricKeyFromBytes([]byte("Nic999Ame888*L-by-5gMurilo-ar-24"))
	if err != nil {
		panic(err)
	}

	return &PasetoToken{
		&token,
		&key,
		&parser,
		duration,
	}, nil
}

func (p PasetoToken) CreateToken(user *domain.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &domain.TokenPayload{
		Id:       id,
		UserId:   user.Id,
		Name:     user.Name,
		UserType: user.UserType,
	}

	err = p.token.Set("payload", payload)
	if err != nil {
		return "", err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(p.duration)

	p.token.SetIssuedAt(issuedAt)
	p.token.SetNotBefore(issuedAt)
	p.token.SetExpiration(expiresAt)

	token := p.token.V4Encrypt(*p.key, nil)

	return token, nil
}

func (p PasetoToken) VerifyToken(token string) (*domain.TokenPayload, error) {
	var payload *domain.TokenPayload

	token = strings.ReplaceAll(token, "Bearer ", "")

	parsedToken, err := p.parser.ParseV4Local(*p.key, token, nil)
	if err != nil {
		return nil, err
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
