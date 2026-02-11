// this service handles all register ,login and similar logic (authentication & password/JWT logic)
package service

import (
	"context"
	"errors"
	"time"

	model "github.com/Rohit-Bhardwaj10/RBAC-Go/internal/models"
	"github.com/Rohit-Bhardwaj10/RBAC-Go/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtSecret  string
	accessttl  time.Duration
	refreshttl time.Duration
}

func NewAuthService(userRepo *repository.UserRepository,
	jwtSecret string,
	accessTTL time.Duration,
	refreshTTL time.Duration) *AuthService {

	return &AuthService{
		userRepo:   userRepo,
		jwtSecret:  jwtSecret,
		accessttl:  accessTTL,
		refreshttl: refreshTTL,
	}
}

func (s *AuthService) Register(ctx context.Context, req *model.User) error {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return errors.New("all fields are required")
	}
	if _, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil {
		return errors.New("invalid creds")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
		RoleID:   req.RoleID,
	}
	if err := s.userRepo.Create(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(ctx context.Context, req *model.User) (*TokenResponse, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid user")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, errors.New("invalid creds")
	}

	return s.generateTokens(user)
}

func (s *AuthService) generateTokens(user *model.User) (*TokenResponse, error) {
	accessClaims := jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(s.accessttl).Unix(),
		"iss":      "rbac-go",
		"type":     "access",
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(s.refreshttl).Unix(),
		"iss":  "rbac-go",
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
	

