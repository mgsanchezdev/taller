package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"workshop-platform/backend/internal/domain"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	jwt.RegisteredClaims
	UserID     string       `json:"user_id"`
	WorkshopID string       `json:"workshop_id"`
	Email      string       `json:"email"`
	Role       domain.Role  `json:"role"`
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: []byte(secret)}
}

type JWT struct {
	secret []byte
}

func (j *JWT) Sign(user *domain.User, expiresIn time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
		UserID:     user.ID,
		WorkshopID: user.WorkshopID,
		Email:      user.Email,
		Role:       user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
