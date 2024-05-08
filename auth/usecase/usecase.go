package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
	"todo-go/auth"
	"todo-go/models"
)

type AutoClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AutoUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAutoUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AutoUseCase {

	return &AutoUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: tokenTTLSeconds * time.Second,
	}
}

func (a *AutoUseCase) SignUp(ctx context.Context, username string, password string) error {
	pwd := sha1.New()

	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AutoUseCase) SignIn(ctx context.Context, username string, password string) (string, error) {
	pwd := sha1.New()

	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)

	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AutoClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AutoUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AutoClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AutoClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
