package users

import (
	"context"
	"time"

	"github.com/google/uuid"
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
		if err := s.repo.UpdateLastLogin(ctx, tx, uid, time.Now()); err != nil {
			tx.Rollback()
			return nil, err
		}
		user.LastLogin = time.Now()
		tx.Commit()
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

	if err := s.repo.Create(ctx, tx, newUser); err != nil {
		tx.Rollback()
		return nil, err
	}

	return newUser, tx.Commit().Error
}

func (s *userService) GetByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error) {
	return s.repo.GetByFirebaseUID(ctx, firebaseUID)
}

// nanti implementasi UpdateProfile di sini
