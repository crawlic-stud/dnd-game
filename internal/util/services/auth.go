package services

import (
	"context"
	"dnd-game/internal/db"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authKey string

type AuthService struct {
	Store          *db.Store
	AuthContextKey authKey
	secret         []byte
	tokenExpires   int
}

func NewAuthService(secret string, tokenExpires int, store *db.Store) *AuthService {
	return &AuthService{
		Store:          store,
		AuthContextKey: authKey("authKey"),
		secret:         []byte(secret),
		tokenExpires:   tokenExpires,
	}
}

// VerifyToken verifies token and returns map containing its data
func (s *AuthService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	user_id := claims["user_id"].(string)
	userUUID, err := uuid.Parse(user_id)
	if err != nil {
		return nil, errors.New("invalid user_id")
	}

	exists, err := s.Store.UserIDExists(context.Background(), userUUID)
	if !exists || err != nil {
		return nil, fmt.Errorf("can't get user with uuid %v", userUUID)
	}

	return claims, nil
}

func (s *AuthService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(s.tokenExpires)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *AuthService) GetUserID(r *http.Request) (uuid.UUID, error) {
	claims, _ := r.Context().Value(s.AuthContextKey).(jwt.MapClaims)
	userID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("cant read user_id")
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cant parse user_id")
	}
	return userUUID, nil
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
