package users

import (
	"context"
	"errors"
	"fmt"
	"gobackend/shared/utils"
	"strings"
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
	user, err := s.repo.GetByUserID(ctx, uid)
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

func (s *userService) GetByUserID(ctx context.Context, firebaseUID string) (*User, error) {
	return s.repo.GetByUserID(ctx, firebaseUID)
}

func (s *userService) AssignRolesToUser(ctx context.Context, userId string, roleIDs []string) (assignRoleRes, error) {
	panic("unimplemented")
}

func (s *userService) GetUsers(ctx context.Context, p utils.Params) ([]GetUsers, utils.Meta, error) {
	users, total, err := s.repo.GetUsers(ctx, p)
	if err != nil {
		return nil, utils.Meta{}, err
	}
	meta := utils.BuildMeta(total, p.Page, p.PerPage)
	return users, meta, nil
}

func (s *userService) CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	if len(req.Roles) == 0 {
		return CreateUserResponse{}, errors.New("roles cannot be empty")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return CreateUserResponse{}, err
	}

	fmt.Println("checking email ...")
	emailExist, err := s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		fmt.Printf("there is something wrong : %f ...", err)
		return CreateUserResponse{}, err
	}
	if emailExist != nil {
		return CreateUserResponse{}, fmt.Errorf("email has already used")
	}

	fmt.Println("ok,email can be used !")
	fmt.Println("starting mapping request ...")

	fid, err := utils.RandID(8)
	if err != nil {
		return CreateUserResponse{}, err
	}

	roles := normalizeRoles(req.Roles)
	u := &User{
		ID:          uuid.NewString(),
		Email:       req.Email,
		Roles:       roles,
		FirebaseUID: fid,
		Password:    string(hashedPw),
		CreatedAt:   time.Now(),
	}

	fmt.Println("mapping clear, inserting data to db ...")

	if err := s.repo.CreateUser(ctx, u); err != nil {
		fmt.Println("failed to insert data , rollback.")
		return CreateUserResponse{}, err
	}

	fmt.Println("data inserted ! returning response to users ")
	return CreateUserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Roles:     req.Roles,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}, nil
}

func normalizeRoles(in []string) []string {
	out := make([]string, 0, len(in))
	for _, r := range in {
		for _, p := range strings.Split(r, ",") {
			v := strings.TrimSpace(p)
			if v != "" {
				out = append(out, v)
			}
		}
	}
	seen := map[string]struct{}{}
	res := make([]string, 0, len(out))
	for _, v := range out {
		k := strings.ToLower(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		res = append(res, v)
	}
	return res
}
