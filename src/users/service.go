package users

import (
	"context"
	"errors"
	"fmt"
	"gobackend/shared/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	repo URepository
	db   *gorm.DB
}

func NewUserService(repo URepository, db *gorm.DB) UService {
	return &userService{repo, db}
}

func (s *userService) SyncFirebaseUser(ctx context.Context, uid, email, name, avatarURL string) (*User, error) {
	user, err := s.repo.GetByFirebaseUID(ctx, uid)
	if err != nil {
		return nil, err
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

	if user != nil {
		err := s.repo.UpdateLastLogin(ctx, tx, uid, time.Now())
		if err != nil {
			return nil, err
		}
		user.LastLogin = time.Now()
		return user, nil
	}

	newUser := &User{
		ID:          uuid.NewString(),
		FirebaseUID: uid,
		Email:       email,
		Name:        name,
		AvatarURL:   avatarURL,
		CreatedAt:   time.Now(),
		LastLogin:   time.Now(),
	}

	err = s.repo.Create(ctx, tx, newUser)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	now := time.Now()
	user, err := s.repo.CheckEmailExist(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
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

	err = s.repo.UpdateLastLogin(ctx, tx, user.FirebaseUID, now)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	accessToken, err := utils.GenerateJWT(user.FirebaseUID)
	if err != nil {
		return nil, err
	}

	refreshToken := utils.GenerateRandomToken()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	if err = s.repo.StoreRefreshToken(ctx, tx, user.ID, refreshToken, expiresAt); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *userService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	rt, err := s.repo.FindRefreshToken(ctx, refreshToken)

	if err != nil || rt.ExpiresAt.Before(time.Now()) {
		return "", errors.New("invalid or expired refresh token")
	}

	token, err := utils.GenerateJWT(rt.UserID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Logout(ctx context.Context, refreshToken string) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := s.repo.DeleteRefreshToken(ctx, tx, refreshToken)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) GetByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error) {
	user, err := s.repo.GetByFirebaseUID(ctx, firebaseUID)
	fmt.Println("firebase uid === ", firebaseUID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateProfile(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("unauthorize")
	}

	if err := s.repo.UpdateProfile(ctx, id); err != nil {
		return err
	}

	return nil
}
