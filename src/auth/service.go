package auth

import (
	"context"
	"errors"
	"time"

	"gobackend/shared/utils"
	"gobackend/src/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	authRepo Repository
	userRepo users.URepository
	db       *gorm.DB
}

func (s *authService) SyncFirebaseUser(ctx context.Context, uid string, email string, name string, avatarURL string) (*users.User, error) {
	panic("unimplemented")
}

func NewAuthService(authRepo Repository, userRepo users.URepository, db *gorm.DB) Service {
	return &authService{authRepo, userRepo, db}
}

func (s *authService) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	now := time.Now()
	user, err := s.userRepo.CheckEmailExist(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("incorrect email or password")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := s.userRepo.UpdateLastLogin(ctx, tx, user.FirebaseUID, now); err != nil {
		tx.Rollback()
		return nil, err
	}

	accessToken, err := utils.GenerateJWT(user.FirebaseUID)
	if err != nil {
		return nil, err
	}

	refreshToken := utils.GenerateRandomToken()
	expiresAt := now.Add(7 * 24 * time.Hour)
	if err := s.authRepo.StoreRefreshToken(ctx, tx, user.ID, refreshToken, expiresAt); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	rt, err := s.authRepo.FindRefreshToken(ctx, refreshToken)
	if err != nil || rt.ExpiresAt.Before(time.Now()) {
		return "", errors.New("invalid or expired refresh token")
	}
	return utils.GenerateJWT(rt.UserID)
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := s.authRepo.DeleteRefreshToken(ctx, tx, refreshToken); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
